use sled::{Db, IVec};
use serde::{Serialize, Deserialize};
use anyhow::{Result, anyhow};

// 你的文件元数据结构
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct FileMeta {
  pub id: String,      // 唯一标识
  pub filename: String,
  pub size: u64,
  // 其他元数据字段
}

// 定义 Repository trait
#[async_trait::async_trait]
pub trait FileMetaRepository {
  async fn insert(&self, meta: FileMeta) -> Result<()>;
  async fn get(&self, id: &str) -> Result<Option<FileMeta>>;
  async fn delete(&self, id: &str) -> Result<()>;
  // 更多方法按需扩展
}

// sled 实现
pub struct SledFileMetaRepository {
  db: Db,
}

impl SledFileMetaRepository {
  pub fn new(db: Db) -> Self {
    Self { db }
  }
}

// 方便序列化/反序列化
// fn serialize(meta: &FileMeta) -> Result<Vec<u8>> {
//   Ok(bincode::serialize(meta)?)
// }
// 
// fn deserialize(bytes: &[u8]) -> Result<FileMeta> {
//   Ok(bincode::deserialize(bytes)?)
// }

// #[async_trait::async_trait]
// impl FileMetaRepository for SledFileMetaRepository {
//   async fn insert(&self, meta: FileMeta) -> Result<()> {
//     let key = meta.id.as_bytes()
//     Ok()
//     // let value = serialize(&meta)?;
//     // sled 的操作是同步的，这里用 tokio::task::spawn_blocking 包裹避免阻塞异步上下文
//     // tokio::task::spawn_blocking(move || {
//     //   self.db.insert(key, value)?;
//     //   self.db.flush()?; // 确保写盘
//     //   Ok::<_, sled::Error>(())
//     // })
//     //   .await?
//     //   .map_err(|e| anyhow!("sled insert error: {}", e))
//   }
// 
//   async fn get(&self, id: &str) -> Result<Option<FileMeta>> {
//     let key = id.as_bytes();
//     let db = self.db.clone();
//     tokio::task::spawn_blocking(move || {
//       match db.get(key)? {
//         Some(ivec) => {
//           let meta = deserialize(&ivec)?;
//           Ok(Some(meta))
//         }
//         None => Ok(None),
//       }
//     })
//       .await?
//       .map_err(|e| anyhow!("sled get error: {}", e))
//   }
// 
//   async fn delete(&self, id: &str) -> Result<()> {
//     let key = id.as_bytes();
//     let db = self.db.clone();
//     tokio::task::spawn_blocking(move || {
//       db.remove(key)?;
//       db.flush()?;
//       Ok::<_, sled::Error>(())
//     })
//       .await?
//       .map_err(|e| anyhow!("sled delete error: {}", e))
//   }
// }
