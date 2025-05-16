use entry::entry_service::EntryService;
use std::net::SocketAddr;
use tokio::net::TcpListener;

pub struct TcpService {
  pub addr: SocketAddr,
}

#[async_trait::async_trait]
impl EntryService for TcpService {
  async fn start(&self) -> anyhow::Result<()> {
    let listener = TcpListener::bind(&self.addr).await?;
    loop {
      let (stream, addr) = listener.accept().await?;
      tokio::spawn(async move {
        println!("Accepted TCP connection from {addr}");
      });
    }
  }

  fn name(&self) -> &'static str {
    "tcp"
  }
}
