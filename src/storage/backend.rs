#[async_trait::async_trait]
pub trait StorageBackend: Send + Sync {
  async fn put(&self, object_id: &str, data: &[u8]) -> anyhow::Result<()>;
  async fn get(&self, object_id: &str) -> anyhow::Result<Vec<u8>>;
  async fn delete(&self, object_id: &str) -> anyhow::Result<()>;
}
