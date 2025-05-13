# FIPS 205 Stateless Hash-Based Digital Signature Standard

[![CI](https://github.com/JoaoHenrique12/fips205/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/JoaoHenrique12/fips205/actions?query=workflow%3Aci)
[![License: Unlicense](https://img.shields.io/badge/license-Unlicense-blue.svg)](http://unlicense.org/)


This repository intends to implement [FIPS 205](https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.205.pdf)
in GoLang. To do so, kasperdi's implementation of [SPHINCS+](https://github.com/kasperdi/SPHINCSPLUS-golang)
is used as a reference, since this repository is indicated in official documentation of [SPINCHS+](https://sphincs.org/software.html).

Informations about CI can be found on [DEVOPS.md](DEVOPS.md).

## [Lamport](lamport/)

Lamport signatures are not part of FIPS 205, but they were the spark to build WOTS signatures and were added to this repository
as a warmup to handle hash functions and the secure rand function.
