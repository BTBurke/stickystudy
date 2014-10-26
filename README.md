StickyStudy List Maker
===
This is a command line tool to turn lists of Chinese vocab in Excel into a format suitable for importing into the iPhone app StickyStudy Chinese.

Install:

`go get github.com/BTBurke/stickystudy`

Usage:
First, create your wordlists in Excel (or Google Docs) in the following format:

<img src="http://i.imgur.com/FZxJH9h.png">

with characters in the first column, pinyin with numbered tones in the second, and the definition in the third.

If you use the Zhongwen Chrome extension, you might want to try my [fork](http://github.com/BTBurke/zhongwen) that will copy a defintion to the clipboard in that format when you push `c`.  You can then paste into your world list with `ctrl + v`.

Use the stickystudy tool to convert your excel word list into the StickyStudy format.

`stickystudy <file>`

This will create one word list for each tab and place them in your Dropbox StickyStudy folder.

In stickystudy, create a list with the same name and then in settings, click restore.  This will load your list of words.
 

