# mi-band-heart-rate-overlay

<img src="https://i.imgur.com/gpPn2Fo.png">

### Requirements

- Mi Band (should work on all version, tested on Mi Band 4)
- The [release](https://github.com/SuspiciousLookingOwl/mi-band-heart-rate-overlay/releases)
- Your Android Phone
  - [Automate](https://play.google.com/store/apps/details?id=com.llamalab.automate&hl=en&gl=US)
  - [Tools & Mi Band](https://play.google.com/store/apps/details?id=cz.zdenekhorak.mibandtools&hl=en&gl=US) ($3.59)
- Phone and your PC connected to a same network

### Setup

##### PC:

1. Download and run the [release](https://github.com/SuspiciousLookingOwl/mi-band-heart-rate-overlay/releases)

##### Android:

1. Install [Tools & Mi Band](https://play.google.com/store/apps/details?id=cz.zdenekhorak.mibandtools&hl=en&gl=US) and [Automate](https://play.google.com/store/apps/details?id=com.llamalab.automate&hl=en&gl=US) on your phone
2. On Automate app, click `More flows...` -> `Start` and search for `Mi Band Heart Rate Overlay`, and Download it
3. Go back to main page (`Flows`), and tap on `Mi Band Heart Rate Overlay`
4. Tap `Start`, and enter your PC IP address
5. It should start sending heart rate data every few seconds, and you should see it being logged on the console app that you ran on the PC

##### OBS:

1. Add `Browser` source, with URL value: `https://suspiciouslookingowl.github.io/mi-band-heart-rate-overlay/static`
2. Input the width and height value, recommended values are `1200` and `600` respectively

---

This tools is inspired by [gergof/StreamHeartMon](https://github.com/gergof/StreamHeartMon)
