## Nostr client CLI application for cross-platform
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
[![Build](https://github.com/nao1215/honeycomb/actions/workflows/build.yml/badge.svg)](https://github.com/nao1215/honeycomb/actions/workflows/build.yml)
[![MultiPlatformUnitTest](https://github.com/nao1215/honeycomb/actions/workflows/unit_test.yml/badge.svg)](https://github.com/nao1215/honeycomb/actions/workflows/unit_test.yml)
[![reviewdog](https://github.com/nao1215/honeycomb/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/nao1215/honeycomb/actions/workflows/reviewdog.yml)

> [!CAUTION]
> Work in progress. Not ready for production.
> Development just started on May 17, 2024. I plan to spend about one hour per day on the project.

![sample](doc/img/new_sample.gif)

Honeycomb is an application that uses the Nostr Protocol to post messages and view trends from the terminal. It is designed to offer both a Command Line Interface and a Text User Interface.

My main purposes for developing Honeycomb are the following:

- To create a comfortable social network as an alternative to X (formerly Twitter) and BlueSky.
- To serve as a testing ground for trying out new technologies.

I enjoy having a free environment and working in the terminal. Therefore, I intend to develop a user-friendly CLI.

## How to install
**go install**
```shell
go install github.com/nao1215/honeycomb@latest
```

**homebrew**
```shell
brew install nao1215/tap/honeycomb
```

## Supported platforms and requirements
- Linux
- macOS
- Windows
- go 1.21 or later

## How to use

### Log in with an existing account
Honeycomb checks for the presence of a private key (e.g. nsec**) available at `${XDG_CONFIG_HOME}/.config/honeycomb/private_key`. If no private key is found, honeycomb will prompt you to enter your private key.   

Honeycomb validates the private key and stores the correct private key locally. 
The feature to create a new account using Honeycomb is not yet implemented. However, we plan to offer this functionality in the future.


![login](./doc/img/nseckey-input.png)

### View timeline, follow users, and more
Once you have logged in, you can view your timeline, follow users, and more. Now, we have several tabs that are not yet implemented.

![timeline](./doc/img/timeline.png)  
  
![follow](./doc/img/follow.png)

### Post a message
You can post a message by pressing the `p` key. The message will be displayed in the timeline of the users who follow you.

![post](./doc/img/post.png)


### Key bindings
- `Ctrl + C`, `q`, `<ESC>`: Exit
- `TAB` : Switch tabs
- `SHIFT + TAB` : Switch tabs in reverse order
- ↑, ↓ : Move cursor
- `p` : Show post form

## Contributing
First off, thanks for taking the time to contribute! See [CONTRIBUTING.md](./CONTRIBUTING.md) for more information.  Contributions are not only related to development. For example, GitHub Star and [GitHub Sponsor](https://github.com/sponsors/nao1215) motivates me to develop!

**Star History**

[![Star History Chart](https://api.star-history.com/svg?repos=nao1215/honeycomb&type=Date)](https://star-history.com/#nao1215/honeycomb&Date)

## Contact
If you would like to send comments such as "find a bug" or "request for additional features" to the developer, please use one of the following contacts.

- [GitHub Issue](https://github.com/nao1215/honeycomb/issues)
- [Nostr](https://nostter.app/npub1xe406xww9s4vla878wd4h6605wqaka9wdammxvpd5he9nysf2z7q0f7804)

## License
This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.


## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tbody>
    <tr>
      <td align="center" valign="top" width="14.28%"><a href="https://debimate.jp/"><img src="https://avatars.githubusercontent.com/u/22737008?v=4?s=75" width="75px;" alt="CHIKAMATSU Naohiro"/><br /><sub><b>CHIKAMATSU Naohiro</b></sub></a><br /><a href="https://github.com/nao1215/honeycomb/commits?author=nao1215" title="Code">💻</a> <a href="https://github.com/nao1215/honeycomb/commits?author=nao1215" title="Documentation">📖</a></td>
    </tr>
  </tbody>
  <tfoot>
    <tr>
      <td align="center" size="13px" colspan="7">
        <img src="https://raw.githubusercontent.com/all-contributors/all-contributors-cli/1b8533af435da9854653492b1327a23a4dbd0a10/assets/logo-small.svg">
          <a href="https://all-contributors.js.org/docs/en/bot/usage">Add your contributions</a>
        </img>
      </td>
    </tr>
  </tfoot>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!
