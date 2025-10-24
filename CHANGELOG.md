# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

<!-- 

## [Unreleased]
## [1.1.11] - DATE
### Added
- 

### Changed
- Tambah endpoint untuk dashboard

### Fixed
- 

### Security
- fix admin auth, hanya system_administrator CRUD admin

-->
## [1.1.10] - DATE
### Fixed
- Get Jadwal Ibadah hari ini menampilkan hanya ibadah yang ada hari ini

## [1.1.9] - 2025-20-24
### Fixed
- Perbedaan jam server dan admin dashboard pada tampilan jadwal ibadah di dashboard
- role error, admin tidak bisa tambah jadwal ibadah

## [1.1.8] - 2025-10-24
### Added
- Tambah endpoint untuk dashboard

### Fixed
- fix admin auth, hanya system_administrator CRUD admin

## [1.1.7] - 2025-10-21
### Added
- crud for admin dashboard endpoint

## [1.1.6] - 2025-10-12
### Added
- CRUD untuk Jadwal Pelayan Altar

## [1.1.5] - 2025-10-10
### Added
- CRUD Finance & Events
- Role Base 

## [1.1.4] - 2025-10-07
### Added
- Seed Default Admin

## [1.1.3] - 2025-10-07
### Added
- Admin login

## [1.1.2] - 2025-10-04
### Added
- endpoint upload warta
- endpoint get latest warta

## [1.1.1] - 2025-10-01
### Added
- CRUD Member untuk Narwastu Connect Dashboard
- Upload Photo Profile Member (supabase)


## [1.0.0] - 2025-09-29
### Added
- 

### Changed
- Optimasi query Postgres untuk perhitungan minggu berjalan

### Fixed
- Perbaikan fungsi untuk endpoint weeklybirthday `response null pada minggu berjalan dengan range akhir bulan dan awal bulan`