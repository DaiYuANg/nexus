use std::path::PathBuf;
use redb::Database;
use tokio::fs;

struct ObjectStorage {
  root_dir: PathBuf,
  meta_db: Database,
}

impl ObjectStorage {
  // 上传对象
  async fn put_object(&self, bucket: &str, key: &str, data: &[u8]) -> anyhow::Result<()> {
    let bucket_path = self.root_dir.join(bucket);
    fs::create_dir_all(&bucket_path).await?;

    let object_path = bucket_path.join(key);
    fs::write(&object_path, data).await?;

    // // 存储元数据
    // self.meta_db.insert(
    //   format!("{}/{}", bucket, key),
    //   data.len().to_be_bytes(),
    // )?;
    // self.meta_db.flush()?;

    Ok(())
  }

  // 下载对象
  async fn get_object(&self, bucket: &str, key: &str) -> anyhow::Result<Vec<u8>> {
    let object_path = self.root_dir.join(bucket).join(key);
    let data = fs::read(object_path).await?;
    Ok(data)
  }

  // 删除对象
  async fn delete_object(&self, bucket: &str, key: &str) -> anyhow::Result<()> {
    let object_path = self.root_dir.join(bucket).join(key);
    fs::remove_file(&object_path).await?;

    // self.meta_db.remove(format!("{}/{}", bucket, key))?;
    // self.meta_db.flush()?;
    Ok(())
  }
}