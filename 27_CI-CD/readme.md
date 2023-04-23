# CI/CD

## CI
- Continuous Integration
- Otomatisasi proses integrasi kode dari beberapa sumber
- misal dari beberapa developer ke satu codebase
- proses:
    1. perubahan kode terjadi
    2. melakukan fetch perubahan
    3. build aplikasi versi terbaru
    4. testing aplikasi
    5. sukses atau gagal proses testing
    6. memberi notifikasi sukses/gagalnya proses testing

## CD
- Continuous Delivery/Deployment
- Otomatisasi proses deployment aplikasi yang telah diverifikasi
- proses:
    1. Unit test
    2. Platform test
    3. Deliver to staging
    4. Application acceptance test
    5. Deploy to prod (manual untuk CDelivery, otomatis untuk CDeployment)
    6. Post deploy test

## Tools CI/CD
tools beragam, yang digunakan dalam projek ini adalah
- code dan commit: git (version control), github (remote repo)
- build dan config: aws dan docker
- test: go test dan github action
- deploy: docker, docker-compose, github action

## Repo tugas
[fork repo](https://github.com/ThiccPan/alta-batch-4-cicd)