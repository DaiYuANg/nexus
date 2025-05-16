use std::net::SocketAddr;
use std::sync::Arc;
use axum::Router;
use axum::routing::get;
use crate::indexer::indexer::IndexService;

mod storage_strategy;
mod indexer;
mod repository;

fn init(){
  // initialize tracing
  tracing_subscriber::fmt::init();
  // dotenvy::dotenv_override().unwrap();
}

#[tokio::main]
async fn main() {
  init();
  // 创建索引服务实例
  let index_service = Arc::new(IndexService::new());

  // 克隆引用，传给后台任务
  let index_service_task = index_service.clone();

  // 启动索引后台任务
  let index_handle = tokio::spawn(async move {
    index_service_task.run().await;
  });

  // 构建简单的 axum 路由
  let app = Router::new().route("/", get(|| async { "Hello, file db!" }));

  // 绑定地址
  let addr = SocketAddr::from(([127, 0, 0, 1], 3000));
  println!("HTTP Listen {}", addr);

  // 启动 HTTP 服务器任务 (这是一个永远运行的 Future)
  let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
  let server_handle = axum::serve(listener, app);

  // 这里演示等待两个任务，你可以改为响应信号优雅退出等
  tokio::select! {
        _ = index_handle => {
            println!("Indexer shutdown");
        }
        _ = server_handle => {
            println!("HTTP shutdown");
        }
    }

  println!("主进程退出");
}
