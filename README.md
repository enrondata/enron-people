# Enron People

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/enrondata/enron-people/workflows/test/badge.svg?branch=master
 [build-status-url]: https://github.com/enrondata/enron-people/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/enrondata/enron-people
 [goreport-url]: https://goreportcard.com/report/github.com/enrondata/enron-people
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/enrondata/enron-people
 [docs-godoc-url]: https://pkg.go.dev/github.com/enrondata/enron-people
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/enrondata/enron-people/blob/master/LICENSE

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

## Source Data

This package is designed to consolidate the information in several data files and eventually output replacements to those files:

* [edo_enron-custodians-data.html](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.html)
* [edo_enron-custodians-data.json](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.json)
* [edo_enron-custodians-data.tsv](https://github.com/enrondata/enrondata/blob/master/data/misc/edo_enron-custodians-data.tsv)

## Links

* [Discussion Tracking in Enron Email using PARAFAC, 2007-04-28](https://www.osti.gov/servlets/purl/1147937)
* [Temporal Text Analysis of Enron Email using Non-negative PARAFAC, 2006-11-01](https://www.osti.gov/servlets/purl/1724662)

