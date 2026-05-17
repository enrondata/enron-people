# Enron People

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/enrondata/enron-people/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/enrondata/enron-people/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/enrondata/enron-people/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/enrondata/enron-people/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/enrondata/enron-people/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/enrondata/enron-people/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/enrondata/enron-people
 [goreport-url]: https://goreportcard.com/report/github.com/enrondata/enron-people
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/enrondata/enron-people
 [docs-godoc-url]: https://pkg.go.dev/github.com/enrondata/enron-people
 [viz-svg]: https://img.shields.io/badge/visualization-Go-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=enrondata%2Fenron-people
 [loc-svg]: https://tokei.rs/b1/github/enrondata/enron-people
 [repo-url]: https://github.com/enrondata/enron-people
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/enrondata/enron-people/blob/main/LICENSE

This package is currently incomplete and a work in progress.

## Overview

This is a package containing metadata on the custodians and people referenced in the Enron Email Dataset. Included information includes:

* Name
* Enron Title
* Enron Email Address
* Custodian Username associated with FERC email distribution

The data is presented in the [System for Cross-domain Identity Management (SCIM) standard format](https://en.wikipedia.org/wiki/System_for_Cross-domain_Identity_Management). User objects follow the [SCIM Core Schema](https://tools.ietf.org/html/rfc7643).

## Usage

This is designed to be used with the FERC Enron Email Dataset as distributed in several forms.

* [CMU Distribution](https://www.cs.cmu.edu/~enron/)
* EDRM Distribution
* FERC Distribution (original)

## Data Files

This repository contains structured data about Enron personnel:

| File | Description | Count |
|------|-------------|-------|
| [data/edo_enron-custodians-data.json](data/edo_enron-custodians-data.json) | Custodians with maildir archives | ~150 |
| data/edo_enron-employees-data.json | All employees found in emails | ~30,000 (planned) |

### Custodians

Custodians are individuals whose email archives are included in the FERC Enron Email Dataset. Each custodian has a corresponding maildir folder in the CMU distribution.

### Employees

The employees file (planned) will contain all unique individuals found in email headers (From, To, Cc, Bcc) across all custodian mailboxes.

## Data Format

Data files use [Schema.org Person](https://schema.org/Person) format:

```json
{
  "@context": "http://schema.org",
  "@type": "Person",
  "alternateName": "allen-p",
  "givenName": "Philip",
  "familyName": "Allen",
  "email": "philip.allen@enron.com",
  "jobTitle": "VP West Desk Gas Trading",
  "maildirs": ["allen-p"]
}
```

## Historical Source Data

This package consolidates information from several original data files:

* [edo_enron-custodians-data.html](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.html)
* [edo_enron-custodians-data.json](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.json)
* [edo_enron-custodians-data.tsv](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.tsv)

## Links

* [Discussion Tracking in Enron Email using PARAFAC, 2007-04-28](https://www.osti.gov/servlets/purl/1147937)
* [Temporal Text Analysis of Enron Email using Non-negative PARAFAC, 2006-11-01](https://www.osti.gov/servlets/purl/1724662)

