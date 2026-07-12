---
title: Viscosity
date: 2026-07-12
slug: v-viscosity
summary: "Viscosity is the resistance of data to movement. Data locked in legacy systems, undocumented schemas, departed engineers. The harder it is to extract data, the higher the viscosity. Viscosity is technical debt in the data supply chain."
tags: data-engineering, viscosity, legacy, extraction, api-design
series: the-vs-of-data
part: 8
---

> Viscosity is not about how far the data must travel. It is about how hard it is to get it moving. Data that is easy to extract is an asset. Data that resists extraction is a hostage. The ransom is engineering time.

Viscosity is the resistance of data to movement through the system. High viscosity: the data is locked in a legacy system with no API, the schema is undocumented, the original engineers have left, and every extraction requires a custom script maintained by the one person who understands the format. Low viscosity: the data is available through a standard API, the schema is published in a catalog, anyone with credentials can query it. The difference between high and low viscosity is the difference between a data platform that can add sources in days and one that adds sources in months.

A bank migrating customer data from a 30-year-old mainframe system illustrates the extreme. The mainframe stores data in VSAM files with COBOL copybooks defining the schema. The copybooks are the only documentation. The COBOL programmers who wrote them have retired. The data must be extracted, parsed according to copybook definitions, transformed to a modern schema, and loaded into the data platform. The extraction alone takes 18 months. The 18 months is the viscosity. The viscosity is the cost.

Viscosity is technical debt in the data supply chain. Every system that produces data but was not designed to expose it adds viscosity. The debt accumulates with every legacy system, every proprietary format, every undocumented schema. The interest on the debt is the engineering time spent extracting data that should have been available through an API. Unlike financial debt, data supply chain debt is invisible — it appears on no balance sheet, is tracked by no metric, and is discovered only when someone tries to use the data. The discovery is the moment the debt becomes visible. The visibility is painful.

The engineering response is to treat data extraction as a first-class requirement for any system that produces data. Every application should expose its data through a standard interface — a read replica, a change data capture feed, an API with documented schemas. The interface should be designed before the application ships. Retrofitting extraction onto an existing system is ten times more expensive than designing it in. The design-in is the investment. The retrofit is the debt repayment. The debt repayment is always more expensive than the investment would have been.

*See: Zhamak Dehghani, "Data Mesh: Delivering Data-Driven Value at Scale" (O'Reilly, 2021) — on domain ownership as the organizational response to viscosity. Maxime Beauchemin, "Functional Data Engineering" (2018) — on treating data pipelines as software engineering with API design standards. Martin Fowler, "Patterns of Enterprise Application Architecture" (Addison-Wesley, 2002), Chapter 9, on data access patterns.*
