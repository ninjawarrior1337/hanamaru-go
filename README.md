# Hanamaru
### A rework of [crocs-and-socks](https://github.com/ninjawarrior1337/crocs-and-socks/)
Sidenote btw: The target was for a 6/9 release for the funny number. Then I remembered that
Nozomi's birthday. Then I remembered that [this magazine cover](https://i.redd.it/dnittba9sfm41.jpg) exists 
and immediately aimed for a 6/9 public "release" (basically just turning off private on GitHub). Also, happy birthday Aidan.

## Whats new?
Everything has been rewritten in Go which makes the new project more memory and CPU efficient.
The framework has been written by hand to provide a nicer dev experience and some custom tooling was built 
to implement linking commands to the framework. Also because Go is the base of the language, the entire project can be
put into a single binary with the help of pkger.

## Ok that's cool but what about the bot iself.
### JP Features
New features such as pitch accent diagram generation (ever watched Dogen), improved Jisho lookup
and returning features such as turning a sentance to romaji

### Audio Playback
Currently, there is a loose concept of a playlist and youtube playback which reqires youtube-dl to be installed on the host
There will be improved queue support and skip support later.

### Roboragi
Similar to the bot on reddit, you can use {*anime name*} to invoke an anime info reply from the bot. Uses AniList to get
the data.

### More Image Editing
Favorites such as poorly edited ゴゴゴ make a return plus some new ones such as JPEG

### And much more.

##Installing
Ok, so about installing...there is no way to install yet besides building an executable by doing
```shell script
git clone https://github.com/ninjawarrior1337/hanamaru-go
go build
```
and specifying the tags with ```-tags``` where possible tags are ```jp,ij```

### The Law Of Equivalent Exchange
Some stuff has yet to be implemented and might not ported at all from crocs-and-socks
at least, not as perfectly as it was. Keep that in mind.


