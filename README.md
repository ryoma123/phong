# phong

[![Build Status](https://travis-ci.org/ryoma123/phong.svg?branch=master)](https://travis-ci.org/ryoma123/phong)

CLI tool to inform the region where a phone number is from.

## Installation

To install, use `go get`:

```sh
$ go get github.com/ryoma123/phong/...
```

## Usage

Pass one phone number as an argument.

```sh
$ phong 0363849000
🇯🇵  Japan -- Tokyo

$ phong 092-651-1111
🇯🇵  Japan -- Fukuoka, Fukuoka
```

Default is Japan. Other countries use country calling code or specify the country code using the region option.

```sh
$ phong +41446681800
🇨🇭  Switzerland -- Zurich

$ phong -r US 6502530000
🇺🇸  United States of America -- Mountain View, CA
```

## Commands

### help, h

Display a help message.

### list, l

Show the list of country code (alpha-2).

```sh
$ phong l
Display all 249 lines? (Y or n)
🇦🇲  AM -- Armenia
🇨🇦  CA -- Canada
🇳🇿  NZ -- New Zealand
...
```

## Options

### --region <country code>, -r <country code>  

Pass phone number region. (default: JP)

### --version, -v

Display the version of phong.

## License

[MIT](https://github.com/ryoma123/phong/blob/master/LICENSE)

## Author

[ryoma123](https://github.com/ryoma123)
