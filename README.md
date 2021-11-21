[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">

<h3 align="center">go-loader</h3>

  <p align="center">
    Simple utility to tool to send concurrent HTTP requests to a list of URLs.
    <br />
    <a href="https://github.com/strixeyecom/go-loader"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/strixeyecom/go-loader">View Demo</a>
    ·
    <a href="https://github.com/strixeyecom/go-loader/issues">Report Bug</a>
    ·
    <a href="https://github.com/strixeyecom/go-loader/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
   </li>
    <li><a href="#downloads">Downloads</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->

## About The Project

This project was initially created to help to generate load for StrixEye Agent.

Creating load for a web application is a very common task for developers. At the very beginning of the load testing
process, we were using a popular tool called [gobuster](https://github.com/OJ/gobuster) to send requests to a list of
URLs.

However, this tool was not designed to be used in a production environment. It was designed to be used in a development
and most importantly to simulate an abrupt raise of traffic to a web application. go-loader was created to help to
simulate a more realistic load. Most of the load and behavior of the application is controlled by random distributions
to simulate the real world.

## Features

- Single visitor can wait before sending a new request, similar to a real world scenario
- Multiple visitors can request simultaneously to create a load
- Visitors can have unique headers, cookies, and custom payloads
- Visitors can have different wait time distributions, endpoints and ip addresses.
    - If the web application supports identifying the originating ip address in the application layer, such as
      [X-Forwarded-For](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For) or
      `CF-CONNECTING-IP`; then visitors can have different ip addresses.

### Built With

Thanks to maintainers and communities of the following projects for making development of this CLI easier. Full list of
dependencies can be found in go modules file.

* [Cobra](https://github.com/spf13/cobra)
* [Viper](https://github.com/spf13/viper)

<!-- GETTING STARTED -->

## Getting Started

General information about setting up go-loader locally

### Prerequisites

Required software and installations.

* go-loader has no external dependencies. It will work on all machines, with supported operating systems.

### Installation

#### Tarball

1. Download [latest-release] for your operating system/architecture
2. Unzip binary and place it somewhere in your path
3. Make it executable

#### Build

Other than Go version 1.16.+, StrixEye CLI has no dependencies/requirements.

Basic building process like the following would suffice.

```shell
   go build -o go-loader cmd/loader.go
```

##### Docker

```shell
    docker pull go-loader:latest
```

<!-- ROADMAP -->

<!-- USAGE EXAMPLES -->

## Usage

### Shell

Following command will create 150 concurrent visitors to send requests the default endpoints.

```shell
go-loader run --target-host=target.omer.beer --target-scheme=https -f /usr/share/common.txt --visitors 150
```

### Docker

The docker image comes as a command line utility, meaning you can access all cli commands.

```shell
docker run go-loader --help
```

The image comes with a
built-in [list](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/common.txt)
of endpoints. Following command will create 150 concurrent visitors to send requests the default endpoints.

```shell
docker run go-loader run --target-host=target.omer.beer --target-scheme=https -f /usr/share/common.txt --visitors 150
```

_For more examples, please refer to the [Documentation](https://pkg.go.dev/strixeyecom/go-loader)_

## Roadmap

See the [open issues](https://github.com/strixeyecom/go-loader/issues) for a list of proposed features (and known
issues).



<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to be learned, inspire, and create. Any
contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the Apache License 2.0 License. See `LICENSE` for more information.



<!-- CONTACT -->

## Contact

StrixEye - [@strixeye](https://twitter.com/strixeye) - help@strixeye.com

Project Link: [https://github.com/strixeyecom/go-loader](https://github.com/strixeyecom/go-loader)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/strixeyecom/go-loader.svg?style=for-the-badge

[contributors-url]: https://github.com/strixeyecom/go-loader/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/strixeyecom/go-loader.svg?style=for-the-badge

[forks-url]: https://github.com/strixeyecom/go-loader/network/members

[stars-shield]: https://img.shields.io/github/stars/strixeyecom/go-loader?style=for-the-badge

[stars-url]: https://github.com/strixeyecom/go-loader/stargazers

[issues-shield]: https://img.shields.io/github/issues/strixeyecom/go-loader.svg?style=for-the-badge

[issues-url]: https://github.com/strixeyecom/go-loader/issues

[license-shield]: https://img.shields.io/github/license/strixeyecom/go-loader.svg?style=for-the-badge

[license-url]: https://github.com/strixeyecom/go-loader/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://linkedin.com/in/strixeye

[latest-release]: https://github.com/strixeyecom/go-loader/releases