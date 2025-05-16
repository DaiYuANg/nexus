use async_trait::async_trait;
use anyhow::Result;
#[derive(Debug, Clone)]
pub struct FileMeta {
  pub name: String,
  pub size: u64,
  pub content_type: String,
  pub checksum: String,
  pub filename: String,
}

#[async_trait]
pub trait FileStorage: Send + Sync {
  async fn store(&self, namespace: &str, id: &str, data: &[u8], meta: FileMeta) -> Result<()>;
  async fn retrieve(&self, namespace: &str, id: &str) -> Result<Vec<u8>>;
  async fn delete(&self, namespace: &str, id: &str) -> Result<()>;
}
