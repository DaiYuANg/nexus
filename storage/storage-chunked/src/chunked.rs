// storage-chunked/src/lib.rs
use async_trait::async_trait;
use std::sync::Arc;
use anyhow::Result;
use storage_core::storage::{FileMeta, FileStorage};

pub struct ChunkedStorage<T: FileStorage> {
  pub inner: Arc<T>,
  pub chunk_size: usize,
}

#[async_trait]
impl<T: FileStorage> FileStorage for ChunkedStorage<T> {
  async fn store(&self, namespace: &str, id: &str, data: &[u8], meta: FileMeta) -> Result<()> {
    let total = data.len();
    let mut idx = 0;
    let mut chunk_index = 0;

    while idx < total {
      let end = (idx + self.chunk_size).min(total);
      let chunk = &data[idx..end];
      let chunk_id = format!("{}.chunk{}", id, chunk_index);

      let chunk_meta = FileMeta {
        name: "".to_string(),
        filename: format!("{}_part{}", meta.filename, chunk_index),
        content_type: meta.content_type.clone(),
        size: chunk.len() as u64,
        checksum: "".to_string(),
      };

      self.inner.store(namespace, &chunk_id, chunk, chunk_meta).await?;

      idx = end;
      chunk_index += 1;
    }

    // 存储 chunk 信息（例如总数）
    let meta_data = format!("{}", chunk_index).into_bytes();
    self.inner.store(namespace, &format!("{}.meta", id), &meta_data, FileMeta {
      name: "".to_string(),
      filename: format!("{}.meta", id),
      content_type: "text/plain".to_string(),
      size: meta_data.len() as u64,
      checksum: "".to_string(),
    }).await?;

    Ok(())
  }

  async fn retrieve(&self, namespace: &str, id: &str) -> Result<Vec<u8>> {
    let meta_bytes = self.inner.retrieve(namespace, &format!("{}.meta", id)).await?;
    let chunk_count = std::str::from_utf8(&meta_bytes)?.parse::<usize>()?;

    let mut full_data = Vec::new();
    for i in 0..chunk_count {
      let chunk_id = format!("{}.chunk{}", id, i);
      let chunk = self.inner.retrieve(namespace, &chunk_id).await?;
      full_data.extend_from_slice(&chunk);
    }
    Ok(full_data)
  }

  async fn delete(&self, namespace: &str, id: &str) -> Result<()> {
    let meta_bytes = self.inner.retrieve(namespace, &format!("{}.meta", id)).await?;
    let chunk_count = std::str::from_utf8(&meta_bytes)?.parse::<usize>()?;

    for i in 0..chunk_count {
      let chunk_id = format!("{}.chunk{}", id, i);
      self.inner.delete(namespace, &chunk_id).await?;
    }
    self.inner.delete(namespace, &format!("{}.meta", id)).await?;
    Ok(())
  }
}
