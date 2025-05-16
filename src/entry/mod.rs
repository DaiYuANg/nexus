pub mod http;
pub mod tcp;

#[async_trait::async_trait]
pub trait EntryService: Send + Sync {
  async fn start(&self) -> anyhow::Result<()>;
  fn name(&self) -> &'static str;
}