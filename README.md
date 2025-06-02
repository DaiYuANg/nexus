# 🚀 MaxIO: Next-Gen High-Performance Object Storage

> 🌐 HTTP / TCP protocol support | 💾 Fast small file upload | 🧠 Auto ID generation | 📁 Virtual directory structure | 🛠️ Embedded-ready

---

## 📝 Overview

**MaxIO** is a high-performance object storage middleware inspired by modern file systems like ZFS and database design patterns. Its goal is to **surpass MinIO**, especially in small file upload performance, while maintaining flexibility, scalability, and embeddability.

---

## ✨ Key Features

- **High-performance small file handling**
- **Write-back memory cache + batch disk flush**
- **Auto-generated primary keys (UUID, Snowflake, or Incremental)**
- **Virtual directory support for object organization**
- **Dual-protocol support: HTTP & TCP**
- **Pluggable metadata backends (e.g. Badger for metadata, BoltDB for config)**

---

## 🧱 Architecture


## Arch
```text
                        +------------------------+
                        |   Client SDK / CLI     |
                        | (HTTP / TCP / gRPC)    |
                        +-----------+------------+
                                    |
               +--------------------+---------------------+
               |                                          |
        +------+-------+                         +--------+-------+
        |   HTTP API   |                         |   TCP Server   |
        |  (Fiber) |                         |  (自定义协议)   |
        +------+-------+                         +--------+-------+
               |                                          |
               +--------------------+---------------------+
                                    |
                           +--------v--------+
                           | Application Core |
                           |------------------|
                           | - 路由分发         |
                           | - 上传/下载        |
                           | - 目录管理        |
                           | - ID生成器        |
                           | - 鉴权&命名空间    |
                           +--------+---------+
                                    |
              +---------------------+----------------------+
              |                                            |
    +---------v----------+                     +-----------v-----------+
    |   Metadata Engine  |                     |    Chunk Engine       |
    |    (BadgerDB)      |                     |     (分片存储)        |
    |--------------------|                     |-----------------------|
    | - Object 元数据     |                     | - 文件分片与合并      |
    | - 虚拟目录结构      |                     | - 热数据 / GC         |
    | - 标签 & 策略       |                     | - Chunk HashMap Index |
    +---------+----------+                     +-----------+-----------+
              |                                            |
     +--------v---------+                         +--------v---------+
     |     bbolt        |                         |   文件存储后端     |
     |------------------|                         |------------------|
     | - 用户/权限/RBAC  |                         | - 本地 FS / S3    |
     | - Bucket 管理     |                         | - 可扩展存储层    |
     +------------------+                         +------------------+
```

## 🛠 Storage Design

- `bbolt` handles:
  - Permissions
  - Buckets
  - Configuration
- `BadgerDB` stores:
  - File metadata
  - Indexes
- File chunks written via memory buffer, flushed asynchronously to reduce disk I/O overhead.

---

## 🚧 Development Goals

- [x] Memory-first write cache
- [x] Background batch disk flush
- [ ] Simple access control layer
- [ ] Namespace-aware object management
- [ ] CLI + optional Web UI

## 💡 Design Inspirations

- ZFS ARC/L2ARC caching model
- RocksDB-style delayed flush
- MinIO's simplicity, extended with performance-first goals

---

## 📦 License

MIT