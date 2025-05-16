use std::path::PathBuf;
use crate::storage::backend::StorageBackend;

pub struct ChunkedStorageBackend {
  base_path: PathBuf,
  chunk_size: usize,
}

#[async_trait::async_trait]
impl StorageBackend for ChunkedStorageBackend {
  async fn put(&self, object_id: &str, data: &[u8]) -> anyhow::Result<()> {
    let chunks = data.chunks(self.chunk_size);
    for (i, chunk) in chunks.enumerate() {
      let chunk_path = self.base_path.join(format!("{}_{}", object_id, i));
      tokio::fs::write(chunk_path, chunk).await?;
    }
    Ok(())
  }

  async fn get(&self, object_id: &str) -> anyhow::Result<Vec<u8>> {
    todo!()
  }

  async fn delete(&self, object_id: &str) -> anyhow::Result<()> {
    todo!()
  }
}
