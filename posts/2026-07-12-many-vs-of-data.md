---
title: The many Vs of data
date: 2026-07-12
slug: many-vs-of-data
summary: "Volume, Velocity, Variety, Veracity, Value, Variability, Visualization, Viscosity, Virality, Volatility. The Vs of data started as a Gartner framework and became a vocabulary for describing why data is hard. Each V names a specific engineering challenge. Each V has a real-world example. Each V is a design constraint."
tags: data-engineering, big-data, volume, velocity, variety
---

In 2001, Doug Laney, an analyst at Gartner, wrote a research note describing three dimensions of the growing data management challenge: Volume, Velocity, and Variety. The note was internal. It was not meant to launch a framework. It launched a framework. The "three Vs of big data" became the standard vocabulary for describing why data was getting harder to manage. Over the following two decades, practitioners added more Vs — Veracity, Value, Variability, Visualization, Viscosity, Virality, Volatility — each naming a specific dimension of the challenge. The proliferation of Vs is itself a data point: data is hard in many ways, and naming the ways is the first step to managing them.

This post explains each V. What it means. Why it matters. A real-world example. The engineering constraint it imposes. The framework is not exhaustive. It is useful. Usefulness beats exhaustiveness.

## Volume: how much

Volume is the amount of data. The most obvious dimension. The least interesting. Volume matters because storage costs money and queries take time. But volume alone is not the hard problem. The hard problem is what you do with the volume — how you store it, how you query it, how you decide what to keep and what to discard.

**Example.** A single Boeing 787 generates approximately 500 GB of sensor data per flight. A fleet of 1,000 aircraft flying two flights per day generates 1 PB per day. The data must be ingested, stored, and analyzed to predict maintenance needs before failures occur. The volume is the constraint. The constraint drives architecture: you cannot store all of it at full resolution forever. You must downsample, aggregate, and tier storage — hot data on SSDs for real-time analysis, warm data on object storage for batch analysis, cold data archived or discarded.

**The engineering constraint.** Volume forces decisions about retention, aggregation, and storage tiering. You cannot keep everything. You must choose what to keep, at what resolution, for how long. The choice is economic: the cost of storing data must be justified by the value of the decisions it enables. Most organizations never make this choice explicitly. They keep everything. The storage bill grows. The value of the oldest data approaches zero. The cost is constant.

*See: James Hamilton, "Internet-Scale Storage" (AWS Reinvent, 2014) — on the physics of storage at hyperscale. Alex Petrov, "Database Internals" (O'Reilly, 2019), Chapter 3, on B-Trees and LSM-Trees — the data structures that make volume queryable.*

## Velocity: how fast

Velocity is the speed at which data arrives and the speed at which it must be processed. The dimension that separates batch from streaming. Batch: data arrives in files, processed hourly or daily. Streaming: data arrives in messages, processed in milliseconds. The choice between batch and streaming is the most consequential architectural decision in data engineering.

**Example.** A fraud detection system at a payment processor. Each transaction generates an event. The event must be evaluated within 100 milliseconds — is this transaction fraudulent? The decision requires consulting historical patterns (has this card been used in this location before?), real-time aggregates (how many transactions has this card made in the past hour?), and ML model inference (what is the fraud score?). The velocity constraint forces the architecture: the historical data must be pre-computed and served from an in-memory cache, the real-time aggregates must be maintained by a stream processor, and the model must be served with sub-millisecond latency.

**The engineering constraint.** Velocity forces a choice between batch and streaming, and streaming is harder. Batch systems are forgiving — if a job fails, you restart it. Streaming systems are unforgiving — if a message is missed, the aggregate is wrong. Streaming requires exactly-once semantics, watermark handling for late data, and state management in the face of failure. The complexity of streaming is the price of low latency. The price is worth paying when the decision must be made now.

*See: Tyler Akidau et al., "The Dataflow Model: A Practical Approach to Balancing Correctness, Latency, and Cost in Massive-Scale, Unbounded, Out-of-Order Data Processing" (VLDB, 2015) — the paper that unified batch and streaming under a single model, now the basis for Apache Beam. Jay Kreps, "Questioning the Lambda Architecture" (O'Reilly, 2014) — why maintaining two code paths for batch and streaming is a maintenance disaster. Martin Kleppmann, "Designing Data-Intensive Applications" (O'Reilly, 2017), Chapter 11, on stream processing and the evolution from Lambda to Kappa architectures.*

## Variety: how diverse

Variety is the diversity of data formats, structures, and sources. The dimension that makes data integration hard. Data arrives as relational tables, JSON documents, XML messages, CSV files, Parquet files, Avro records, Protobuf messages, images, videos, log files, and free-text documents. Each format has its own schema, its own semantics, its own quirks. Integrating them requires understanding each one.

**Example.** A hospital integrates data from electronic health records (structured, HL7 format), lab systems (semi-structured, ASTM format), imaging systems (unstructured, DICOM format), patient surveys (free text), and wearable devices (JSON streams). A single patient's record spans five systems, four formats, and three data models. Integrating them requires mapping each source to a common model, resolving identifier conflicts (is patient 123 in the EHR the same person as patient 456 in the lab system?), and handling the different update cadences of each source.

**The engineering constraint.** Variety forces schema design and data integration. Every new data source requires a pipeline. Every pipeline requires understanding the source schema, mapping it to the target schema, and handling the edge cases. The number of pipelines grows with the number of sources. The maintenance cost grows with the number of pipelines. Variety is the reason data engineering teams grow faster than the data they manage. Each new source adds a pipeline. Each pipeline adds a maintenance burden. The burden is the cost of variety.

*See: Serge Abiteboul et al., "Data on the Web: From Relations to Semistructured Data and XML" (Morgan Kaufmann, 1999) — the classic text on the shift from structured to semi-structured data. Sanjoy Dasgupta and Peter L. Bartlett, "Schema Mapping and Data Exchange" (PODS, 2003) — the theoretical foundations of mapping between schemas. Joe Reis and Matt Housley, "Fundamentals of Data Engineering" (O'Reilly, 2022), Chapter 4, on the modern data integration landscape from Fivetran to dbt.*

## Veracity: how trustworthy

Veracity is the quality and trustworthiness of data. The dimension that determines whether anyone uses the data platform. Data can be inaccurate (wrong values), incomplete (missing values), inconsistent (contradictory values across sources), or ambiguous (values whose meaning is unclear). Veracity is the hardest V because it is not purely technical — it requires domain knowledge to assess.

**Example.** A weather forecasting system ingests data from ground stations, satellites, weather balloons, and citizen reports. Ground stations are accurate but sparse. Satellites cover the globe but have lower resolution. Weather balloons provide vertical profiles but are launched twice daily. Citizen reports are abundant but unreliable — people report hail when they hear acorns on the roof. The system must fuse data from sources of varying veracity, weighting each source by its historical accuracy, and produce a forecast that is more accurate than any single source.

**The engineering constraint.** Veracity forces data quality testing, lineage tracking, and provenance. You must know where each data point came from, how it was transformed, and how reliable it is. You must test for the failures you can predict (nulls in required fields, values outside expected ranges) and monitor for the failures you can't (systematic drift, new failure modes introduced by source system changes). Veracity is the dimension that separates a data platform that produces numbers from a data platform that produces numbers people trust.

*See: Tom Redman, "Data Driven: Profiting from Your Most Important Business Asset" (Harvard Business Review Press, 2008) — the foundational text on treating data quality as a management discipline. Barr Moses et al., "Data Quality Management at Scale" (Monte Carlo, 2022) — modern operational practices for data observability. Wenfei Fan and Floris Geerts, "Foundations of Data Quality Management" (Morgan & Claypool, 2012) — a formal treatment of data consistency, currency, and accuracy constraints.*

## Value: how useful

Value is the economic worth of data. The dimension that justifies the existence of the data platform. Data has value when it informs a decision that produces a better outcome. Data has no value when it sits in a table that nobody queries. Most data in most organizations has no value. It was collected because it could be, stored because storage is cheap, and never used because nobody knew what to do with it.

**Example.** An e-commerce company collects every click, every page view, every add-to-cart, every purchase. The clickstream data is 10 TB per day. The data team builds pipelines to ingest it, transform it, and serve it to analysts. Six months later, nobody is querying the clickstream tables. The analysts are querying the purchase data and ignoring the rest. The clickstream data has zero value. The storage cost is $2,400 per month. The pipeline maintenance cost is one engineer at 20% time. The total cost exceeds the value. The data is a liability.

**The engineering constraint.** Value forces prioritization. You cannot pipeline every data source. You must pipeline the sources that will be used. Predicting which sources will be used requires understanding the business questions that need answering. The understanding requires talking to the people who will use the data. The talking is the most underinvested activity in data engineering. Engineers build pipelines for data sources they have, not for questions the business needs answered. The pipeline exists. The value doesn't.

*See: Doug Laney, "Infonomics: How to Monetize, Manage, and Measure Information as an Asset for Competitive Advantage" (Routledge, 2017) — the book that extended Laney's original Gartner note into a framework for treating data as a balance-sheet asset. Thomas C. Redman, "Data's Credibility Problem" (Harvard Business Review, 2013) — why most corporate data has negative net value. Benn Stancil, "The Data Platform Cost Model" (Mode, 2020) — a practical framework for calculating whether a data pipeline is worth building.*

## Variability: how inconsistent

Variability is the fluctuation in data characteristics over time. The dimension that breaks pipelines. A data source that produces 1,000 events per hour suddenly produces 100,000. A column that has always contained integers suddenly contains strings. A partition that has always had data suddenly is empty. Variability is the reason pipelines need monitoring, alerting, and circuit breakers.

**Example.** A food delivery platform ingests order data from its mobile app. Order volume follows a predictable pattern: peaks at lunch and dinner, troughs in between, higher on weekends. On Super Bowl Sunday, order volume spikes to 20× the normal peak. The ingestion pipeline, sized for 2× headroom, falls behind. The transformation pipeline, expecting 1 million rows, receives 20 million. The dashboard queries time out. The data platform fails at the moment it is most needed.

**The engineering constraint.** Variability forces elastic infrastructure and defensive pipeline design. Pipelines must be sized for peaks, not averages. The peak-to-average ratio determines the cost of infrastructure. A pipeline with a 20:1 peak-to-average ratio costs 20 times more to run on elastic infrastructure than a pipeline with constant load. The cost is the insurance against variability. The insurance premium is paid whether the peak occurs or not.

*See: John Allspaw, "The Art of Capacity Planning" (O'Reilly, 2008) — the operations engineering approach to sizing for variability. Betsy Beyer et al., "Site Reliability Engineering" (O'Reilly, 2016), Chapter 13, on managing overload and load shedding — the SRE playbook for variability. Daniel Abadi et al., "The Design and Implementation of Modern Column-Oriented Database Systems" (Foundations and Trends in Databases, 2013) — why elastic cloud infrastructure changes the design trade-offs for handling variable workloads.*

## Visualization: how understandable

Visualization is the presentation of data in a form that humans can perceive patterns, anomalies, and relationships. The dimension that bridges the gap between data and decision. A table of numbers is data. A chart is information. The difference is the visual encoding — position, length, color, shape — that maps data attributes to perceptual channels. The mapping is the design. The design determines whether the viewer sees the signal or the noise.

**Example.** A hospital's infection control dashboard tracks surgical site infections across 12 operating theaters. The raw data is a table of 50,000 surgeries with 200 columns. The dashboard visualizes the infection rate as a control chart — a line graph with upper and lower control limits. When a theater's infection rate exceeds the upper control limit, the point turns red. The red point is the signal. The signal triggers an investigation. The investigation finds that Theater 7's sterilization protocol was changed three weeks ago. The change caused the spike. The spike was visible because the visualization made the anomaly perceptible.

**The engineering constraint.** Visualization forces data to be aggregated, cleaned, and structured for the chart, not for the database. The data model that serves a dashboard is different from the data model that serves an ad-hoc query. The dashboard model must be pre-computed, pre-aggregated, and optimized for the specific charts it serves. The ad-hoc model must be flexible. The two models diverge. The divergence is the cost of visualization.

*See: Edward Tufte, "The Visual Display of Quantitative Information" (Graphics Press, 1983) — the foundational text on data visualization as cognitive engineering. Jacques Bertin, "Semiology of Graphics" (1967, English translation 1983) — the pre-digital taxonomy of visual variables (position, size, value, texture, color, orientation, shape) that still defines the field. Leland Wilkinson, "The Grammar of Graphics" (Springer, 1999) — the theoretical foundation of ggplot2 and Vega-Lite, formalizing visualization as a composition of graphical primitives.*

## Viscosity: how resistant to flow

Viscosity is the resistance of data to movement through the system. The dimension that describes how hard it is to get data from where it is to where it needs to be. High viscosity: the data is locked in a legacy system with no API, the schema is undocumented, the owner left the company, and every extraction requires a custom script. Low viscosity: the data is available through a standard API, the schema is documented in a catalog, and ingestion is a configuration change.

**Example.** A bank wants to migrate customer data from a 30-year-old mainframe system to a modern data platform. The mainframe stores data in VSAM files with COBOL copybooks defining the schema. The copybooks are the only documentation. The COBOL programmers who wrote them are retired. The data must be extracted, parsed according to the copybook definitions, transformed to a modern schema, and loaded. The extraction takes 18 months. The viscosity is the time. The time is the cost.

**The engineering constraint.** Viscosity forces investment in data extraction tooling and API design. Every system that produces data should expose it through a standard interface. The interface should be documented. The documentation should be machine-readable. The absence of these things is viscosity. Viscosity is technical debt in the data supply chain. The debt compounds. The interest is the engineering time spent extracting data from systems that were not designed to be extracted from.

*See: Zhamak Dehghani, "Data Mesh: Delivering Data-Driven Value at Scale" (O'Reilly, 2021) — on domain ownership as the organizational response to data viscosity. Maxime Beauchemin, "Functional Data Engineering" (2018) — on treating data pipelines as software engineering, with the same standards for API design, testing, and documentation. Martin Fowler, "Patterns of Enterprise Application Architecture" (Addison-Wesley, 2002), Chapter 9, on the repository pattern and data access layers — the pre-big-data patterns that anticipated the viscosity problem.*

## Virality: how fast it spreads

Virality is the rate at which data usage spreads through an organization. The dimension that describes adoption. A dataset that one team finds useful spreads to other teams. The usage grows. The growth is viral because each new user discovers new questions the data can answer, which attracts more users. Virality is the metric that separates successful data platforms from unsuccessful ones.

**Example.** An HR analytics team builds a dataset of employee movement — hires, promotions, transfers, departures. The recruiting team discovers it and starts using it to measure time-to-fill for open positions. The finance team discovers it and starts using it for headcount forecasting. The facilities team discovers it and starts using it for office space planning. The dataset spreads virally because it answers questions that multiple teams have. The virality is evidence that the dataset is valuable.

**The engineering constraint.** Virality forces investment in documentation, discoverability, and data literacy. A dataset that no one knows about cannot spread. A dataset that no one understands cannot spread. A dataset that no one trusts cannot spread. The data catalog is the mechanism for discoverability. The data dictionary is the mechanism for understanding. The data quality dashboard is the mechanism for trust. Without these, virality is zero. The data exists. It is invisible. Invisibility is the enemy of value.

*See: Alation, "The Data Catalog: The Foundation of Data Culture" (Whitepaper, 2020) — on how catalogs enable data discoverability and adoption. Prukalpa Sankar, "The Data Maturity Curve" (Atlan, 2021) — on the stages of data culture from fragmented to viral. Michelle Casbon et al., "Data Governance: The Definitive Guide" (O'Reilly, 2021) — on how governance enables rather than restricts data virality.*

## Volatility: how long it matters

Volatility is the rate at which data loses relevance over time. The dimension that determines retention policy. High volatility: the data is valuable for hours or days, then worthless. Low volatility: the data is valuable for years or decades. Volatility determines how long to store data and how much to invest in its quality.

**Example.** A stock trading system ingests real-time price data. The data is valuable for milliseconds — the time window for executing an arbitrage. After the window closes, the data's value drops to near zero for trading but remains valuable for backtesting and compliance. The dual value curve — immediate and high, then residual and low — determines the storage architecture: in-memory for the trading window, archival for compliance. The volatility is the reason for the two-tier storage strategy.

**The engineering constraint.** Volatility forces tiered storage and retention policies. Data with high short-term value and low long-term value should be stored in fast, expensive storage for a short period, then archived or deleted. Data with sustained value should be stored in cost-optimized storage for the long term. The retention policy should be explicit. The enforcement should be automated. Most organizations have no retention policy. The absence is a policy — keep everything forever. Forever is expensive.

*See: Apache Iceberg, "Table Spec: Snapshots and Time Travel" (Iceberg Documentation) — on the table format features that make retention management programmable. Netflix Tech Blog, "Evolution of Data Lifecycle Management at Netflix" (2021) — a real-world architecture for automated data retention. Gordon Moore, "Cramming More Components onto Integrated Circuits" (Electronics, 1965) — the paper that predicted storage would become cheap enough that keeping everything seemed rational, even when it wasn't.*

## The Vs as a diagnostic tool

The Vs are not a taxonomy. They are a diagnostic tool. When a data project is struggling, ask: which V is the problem? Is the volume overwhelming the storage tier? Is the velocity exceeding the pipeline's throughput? Is the variety creating an unmanageable number of pipelines? Is the veracity undermining trust in the numbers? Is the value absent because nobody knows what questions the data should answer?

The Vs name the dimensions of difficulty. Naming the difficulty is the first step to managing it. The management is engineering. The engineering is the response to the V. The V is the constraint. The constraint is the design.

---

**References by V:**

*Volume* — James Hamilton, "Internet-Scale Storage," AWS Reinvent, 2014. Alex Petrov, *Database Internals*, O'Reilly, 2019.

*Velocity* — Tyler Akidau et al., "The Dataflow Model," VLDB, 2015. Jay Kreps, "Questioning the Lambda Architecture," O'Reilly, 2014. Martin Kleppmann, *Designing Data-Intensive Applications*, O'Reilly, 2017.

*Variety* — Serge Abiteboul et al., *Data on the Web: From Relations to Semistructured Data and XML*, Morgan Kaufmann, 1999. Joe Reis and Matt Housley, *Fundamentals of Data Engineering*, O'Reilly, 2022.

*Veracity* — Tom Redman, *Data Driven: Profiting from Your Most Important Business Asset*, Harvard Business Review Press, 2008. Barr Moses et al., "Data Quality Management at Scale," Monte Carlo, 2022. Wenfei Fan and Floris Geerts, *Foundations of Data Quality Management*, Morgan & Claypool, 2012.

*Value* — Doug Laney, *Infonomics*, Routledge, 2017. Thomas C. Redman, "Data's Credibility Problem," Harvard Business Review, 2013. Benn Stancil, "The Data Platform Cost Model," Mode, 2020.

*Variability* — John Allspaw, *The Art of Capacity Planning*, O'Reilly, 2008. Betsy Beyer et al., *Site Reliability Engineering*, O'Reilly, 2016. Daniel Abadi et al., "The Design and Implementation of Modern Column-Oriented Database Systems," Foundations and Trends in Databases, 2013.

*Visualization* — Edward Tufte, *The Visual Display of Quantitative Information*, Graphics Press, 1983. Jacques Bertin, *Semiology of Graphics*, 1967. Leland Wilkinson, *The Grammar of Graphics*, Springer, 1999.

*Viscosity* — Zhamak Dehghani, *Data Mesh*, O'Reilly, 2021. Maxime Beauchemin, "Functional Data Engineering," 2018. Martin Fowler, *Patterns of Enterprise Application Architecture*, Addison-Wesley, 2002.

*Virality* — Prukalpa Sankar, "The Data Maturity Curve," Atlan, 2021. Michelle Casbon et al., *Data Governance: The Definitive Guide*, O'Reilly, 2021.

*Volatility* — Apache Iceberg, "Table Spec: Snapshots and Time Travel." Netflix Tech Blog, "Evolution of Data Lifecycle Management at Netflix," 2021. Gordon Moore, "Cramming More Components onto Integrated Circuits," Electronics, 1965.

**Foundational:**
- Doug Laney, "3D Data Management: Controlling Data Volume, Velocity, and Variety," Gartner, 2001.
- Related posts: [Data engineering series](https://blog.hackspree.com/#what-is-data-engineering), [The economics of data](https://blog.hackspree.com/#data-engineering-economics)
