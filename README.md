<img src="https://static1.pixcee.dev/external/strawstack/logo.png" width="128px"></img>

# Strawhouse
Bare minimal file store engine, featured lightweight indexing, signed checksum file integrity, and pre-signed token with access control.

## Architecture
- Store file using host filesystem-backed without overhead.
- Supports signed checksums and flags (file-level validation for compromised direct filesystem modifications) using [xattr](https://en.wikipedia.org/wiki/Extended_file_attributes) stored directly in [inode](https://en.wikipedia.org/wiki/Inode) for lightning fast access and no database overhead.
- Reengineered stateless pre-signed token for clients to upload and retrieve file with expiration time, path restriction, and custom attribute support. Used only 27-bytes token with bit-by-bit optimization (see [performance benchmark](https://github.com/strawst/strawhouse/wiki/Benchmark)).
- Get file directly from HTTP GET (with just path and token), upload using pre-defined API endpoints.
- Store file metadata in high performance key-value embed database: [pogreb](https://github.com/akrylysov/pogreb).
- Designed and optimized for store and serve millions of static files from the ground up.
- Structured to support file validation, file event hooks in the future.

## Module

Get into each modules for more documentation:

| [Backend](#backend)                       | [Driver](#driver)                        | [Command](#command)                         |
|-------------------------------------------|------------------------------------------|---------------------------------------------|
| File server with HTTP and gRPC interface. | Go library for interacting with backend. | Command line tool to test and manage files. |

## Documentation

To learn more about Strawhouse, see the [documentation](https://strawhouse.doc.pixcee.dev/).

## Credits

Developed by **[BSthun](https://github.com/BSthun)**, originated to use internally at [Connected Tech](https://www.connectedtech.co.th) and [Pixcee](https://www.pixcee.app/).

The name of **Straw House** inspired by pronunciation of "Store House" (in the meaning of file store), plus [Facebook's Haystack](https://engineering.fb.com/2009/04/30/core-infra/needle-in-a-haystack-efficient-storage-of-billions-of-photos/) and a song [Jade - Straw House](https://open.spotify.com/track/50uwQoov3D7ASWwfmRVHQI?si=9081a42990ba4233).