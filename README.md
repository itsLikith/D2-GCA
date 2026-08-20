<div align="center">
  <img src="/docs/assets/banner.png" alt="PRISM Banner" width="100%">
  
  <h1>PRISM</h1>
  <h3>Positioning & Resilience through Intelligent Signal Mapping</h3>

  <p>
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg?style=for-the-badge" alt="License">
    <img src="https://img.shields.io/badge/Status-Active-success.svg?style=for-the-badge" alt="Status">
    <img src="https://img.shields.io/badge/Domain-Aviation%20Navigation-0A66C2.svg?style=for-the-badge" alt="Domain">
  </p>
</div>

<br>

## Overview

The aviation sector depends on continuous and accurate navigation data to ensure safe flight operations. During GPS outages, aircraft face significant challenges in determining their precise location. Traditional backup navigation methods are limited and may not always provide reliable positioning.

**PRISM** addresses this critical challenge by providing a mathematical and machine-assisted positioning model that calculates an aircraft’s location using signals from two nearby Distance Measuring Equipment (DME) stations. By combining geometric principles with trained computational methods, PRISM enables robust, GPS-independent position estimation.

This approach improves navigation resilience, enhances situational awareness, and supports safe and efficient aircraft operations in GPS-denied environments.

---

## Key Features

| Capability | Description |
| :--- | :--- |
| **GPS-Independent Positioning** | Reliable location estimation without reliance on satellite signals |
| **Dual-DME Geometry** | Uses range measurements from two DME stations for accurate triangulation |
| **Mathematical + Machine-Assisted Model** | Combines classical geometric principles with computational intelligence |
| **Enhanced Situational Awareness** | Provides continuous position data during GPS outages |
| **Resilient Navigation Support** | Designed for real-world aviation operational requirements |

---

## How It Works

PRISM determines aircraft position through a structured multi-stage process:

1. **Signal Acquisition** — Receives range (slant-range) measurements from two nearby DME ground stations  
2. **Geometric Solution** — Applies geometric intersection principles to solve for possible aircraft locations  
3. **Computational Refinement** — Uses trained methods to refine the solution, resolve ambiguities, and improve accuracy  
4. **Position Output** — Delivers a reliable position estimate suitable for navigation support in GPS-denied conditions  

---

## Motivation

GPS vulnerability remains a significant concern in modern aviation. PRISM offers a practical, ground-based alternative that leverages existing DME infrastructure to maintain navigation capability when satellite-based systems are unavailable or degraded.

---

## Project Status

**Active Development**  
Core positioning model and validation framework are under continuous improvement.

---

## Contributing

Contributions, issues, and feature requests are welcome.  
Please check the [issues page](../../issues) if you would like to contribute.

---

## License

This project is licensed under the **Apache License 2.0**.  
See the [LICENSE](LICENSE) file for full details.

<br>

<div align="center">
  <sub>
    <strong>PRISM</strong> — Resilient Positioning for Modern Aviation
  </sub>
</div>