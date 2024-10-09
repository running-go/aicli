# aicli
AICLI is an AI-powered CLI tool built with Golang that answers your questions using the Google Gemini API or Other API More.


## 🍎 Getter Started
To get started with AICLI, you can install it using the following method:
### Installation
To install the CLI, use the command below:
```bash
go install github.com/running-go/aicli@latest
```
Go will automatically install it in your $GOPATH/bin directory, which should be in your $PATH.


### Usage aicli gemini
Once installed, you can use the aicli gemini CLI command. To confirm installation, type aicli gemini at the command line.

AICLI gemini uses the Google Gemini API, so you need to set the API key. To get the API key (It's FREE), visit here and set it in the environment variable GEMINI_API_KEY:

```bash
export GEMINI_API_KEY=<API_KEY>
```
The above method sets the API key for the current session only. To set it permanently, add the above line to your `.bashrc` or `.zshrc` file.

> **Note:** If you encounter the error command not found: aicli gemini, you need to add `$GOPATH/bin` to your `$PATH` environment variable. For more details, refer to [this guide](https://gist.github.com/Pradumnasaraf/ca6f9a0507089a4c44881446cdda4aa3).


### Commands
```text
(base) ➜  aicli git:(feat/cha_v1) ✗ aicli gemini --help                                          
A CLI tool to interact with the Gemini API. You can ask questions and get responses in text or image format.

Usage:
  aicli gemini [flags]
  aicli gemini [command]

Available Commands:
  image       Know details about an image (Please put your question in quotes)
  search      Ask a question and get a response (Please put your question in quotes)

Flags:
  -h, --help   help for gemini

Use "aicli gemini [command] --help" for more information about a command.

```

>  eg: aicli gemini search "What is kubernetes" --words 525

>  eg: aicli gemini image "What is this image about?" --path /path/to/image.jpg --format jpg

### 📜 License

This project is licensed under the Apache-2.0 license - see the [LICENSE](LICENSE) file for details.