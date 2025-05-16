// src/entry/tcp.rs
use super::EntryService;
use tokio::net::TcpListener;
use std::net::SocketAddr;
use crate::app_state::AppState;

pub struct TcpService {
  pub addr: SocketAddr,
  pub app_state: AppState,
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
