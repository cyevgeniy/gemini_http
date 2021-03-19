# gemini_http
Simple client for viewing gemini files via HTTP.

## Why

Gemini protocol is awesome, but it is also too hard to leave http to start using
new protocol(personaly I think that new protocol was unneccasary(but it is still awesome) -
using just gemini markup language is already does the same job, but simpler).
```Gemini_http``` allows to
use gemini files, but still staying at http(s) protocol - just drop
your files somewhere at your server, and that's all.

## Usage

* ```o``` - Open specified url (For example ```https://somedomain.com/index.gmi```).
* ```l``` - View links from current page
* ```lX```(X - link number) - Open link
* ```h``` - View url history
* ```hX```(X - url number) - Open previously visited url
* ```q``` - quit

### On Windows

On windows, ```gemini_http``` works via ```git-bash.exe```. It is the terminal emulator
that comes with git for windows. If you have git installed, you probably
already has ```git-bash```.

### Proxy

If you want to use proxy, set is with command:

```bash
export HTTP_PROXY="http://proxyIp:proxyPort"
```

After that, gemini_http will use this proxy to send requests.

## Some screenshots

![Screenshot 1](https://github.com/cyevgeniy/gemini_http/blob/master/scr1.png)
![Screenshot 2](https://github.com/cyevgeniy/gemini_http/blob/master/scr2.png)
