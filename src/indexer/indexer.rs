use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Notify;
use tokio::time::sleep;

pub struct IndexService {
  shutdown_notify: Arc<Notify>,
}

impl IndexService {
  pub(crate) fn new() -> Self {
    Self {
      shutdown_notify: Arc::new(Notify::new()),
    }
  }

  pub(crate) async fn run(&self) {
    println!("indexer running");

    loop {
      tokio::select! {
          _ = self.shutdown_notify.notified() => {
              println!("索引服务收到关闭信号，准备退出");
              break;
          }
          _ = sleep(Duration::from_secs(5)) => {
              println!("索引服务执行周期性任务，比如构建/维护索引");
              // 这里写索引处理逻辑
          }
      }
    }

    println!("索引服务退出");
  }

  fn shutdown(&self) {
    self.shutdown_notify.notify_waiters();
  }
}
