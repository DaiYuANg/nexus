use std::path::PathBuf;
use crate::storage::backend::StorageBackend;

pub struct FileStorageBackend {
  base_path: PathBuf,
}

#[async_trait::async_trait]
impl StorageBackend for FileStorageBackend {
  async fn put(&self, object_id: &str, data: &[u8]) -> anyhow::Result<()> {
    let path = self.base_path.join(object_id);
    tokio::fs::write(path, data).await?;
    Ok(())
  }

  async fn get(&self, object_id: &str) -> anyhow::Result<Vec<u8>> {
    todo!()
  }

  async fn delete(&self, object_id: &str) -> anyhow::Result<()> {
    todo!()
  }
}