# Spinners

Go implementation of the [sindresorhus/cli-spinners](https://github.com/sindresorhus/cli-spinners).

## Usage

```go
# create new dots spinner
spinner := spinners.NewDotsSpinner(os.Stdout)

# start spinner in a separate go routing
spinner.Start("waiting for something long ...")

# do something
...

# stop spinner and reset the line
spinner.Stop()

# or get a spinner by its name
spinner, err := spinners.NewSpinnerByName("dots", os.Stdout)
if err != nil {
    log.Fatal(err)
}
```

## Test a Spinner

```bash
go run example/app.go nameofspinner
```

## Build

The spinners.go can be regenerated using the `go generate` command. The command will load the latest spinners.json from
master branch of [sindresorhus/cli-spinners](https://github.com/sindresorhus/cli-spinners) and regenerate the
spinners.go and the README.md file.

## Available Spinners
| Name | Interval | Frames |
| ---- | -------- | ------ |
| arc | 100 | ◜ ◠ ◝ ◞ ◡ ◟  |
| arrow | 100 | ← ↖ ↑ ↗ → ↘ ↓ ↙  |
| arrow2 | 80 | ⬆️  ↗️  ➡️  ↘️  ⬇️  ↙️  ⬅️  ↖️   |
| arrow3 | 120 | ▹▹▹▹▹ ▸▹▹▹▹ ▹▸▹▹▹ ▹▹▸▹▹ ▹▹▹▸▹ ▹▹▹▹▸  |
| balloon | 140 |   . o O @ *    |
| balloon2 | 120 | . o O ° O o .  |
| bounce | 120 | ⠁ ⠂ ⠄ ⠂  |
| bouncingBall | 80 | ( ●    ) (  ●   ) (   ●  ) (    ● ) (     ●) (    ● ) (   ●  ) (  ●   ) ( ●    ) (●     )  |
| bouncingBar | 80 | [    ] [=   ] [==  ] [=== ] [ ===] [  ==] [   =] [    ] [   =] [  ==] [ ===] [====] [=== ] [==  ] [=   ]  |
| boxBounce | 120 | ▖ ▘ ▝ ▗  |
| boxBounce2 | 100 | ▌ ▀ ▐ ▄  |
| christmas | 400 | 🌲 🎄  |
| circle | 120 | ◡ ⊙ ◠  |
| circleHalves | 50 | ◐ ◓ ◑ ◒  |
| circleQuarters | 120 | ◴ ◷ ◶ ◵  |
| clock | 100 | 🕐  🕑  🕒  🕓  🕔  🕕  🕖  🕗  🕘  🕙  🕚   |
| dots | 80 | ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏  |
| dots10 | 80 | ⢄ ⢂ ⢁ ⡁ ⡈ ⡐ ⡠  |
| dots11 | 100 | ⠁ ⠂ ⠄ ⡀ ⢀ ⠠ ⠐ ⠈  |
| dots12 | 80 | ⢀⠀ ⡀⠀ ⠄⠀ ⢂⠀ ⡂⠀ ⠅⠀ ⢃⠀ ⡃⠀ ⠍⠀ ⢋⠀ ⡋⠀ ⠍⠁ ⢋⠁ ⡋⠁ ⠍⠉ ⠋⠉ ⠋⠉ ⠉⠙ ⠉⠙ ⠉⠩ ⠈⢙ ⠈⡙ ⢈⠩ ⡀⢙ ⠄⡙ ⢂⠩ ⡂⢘ ⠅⡘ ⢃⠨ ⡃⢐ ⠍⡐ ⢋⠠ ⡋⢀ ⠍⡁ ⢋⠁ ⡋⠁ ⠍⠉ ⠋⠉ ⠋⠉ ⠉⠙ ⠉⠙ ⠉⠩ ⠈⢙ ⠈⡙ ⠈⠩ ⠀⢙ ⠀⡙ ⠀⠩ ⠀⢘ ⠀⡘ ⠀⠨ ⠀⢐ ⠀⡐ ⠀⠠ ⠀⢀ ⠀⡀  |
| dots2 | 80 | ⣾ ⣽ ⣻ ⢿ ⡿ ⣟ ⣯ ⣷  |
| dots3 | 80 | ⠋ ⠙ ⠚ ⠞ ⠖ ⠦ ⠴ ⠲ ⠳ ⠓  |
| dots4 | 80 | ⠄ ⠆ ⠇ ⠋ ⠙ ⠸ ⠰ ⠠ ⠰ ⠸ ⠙ ⠋ ⠇ ⠆  |
| dots5 | 80 | ⠋ ⠙ ⠚ ⠒ ⠂ ⠂ ⠒ ⠲ ⠴ ⠦ ⠖ ⠒ ⠐ ⠐ ⠒ ⠓ ⠋  |
| dots6 | 80 | ⠁ ⠉ ⠙ ⠚ ⠒ ⠂ ⠂ ⠒ ⠲ ⠴ ⠤ ⠄ ⠄ ⠤ ⠴ ⠲ ⠒ ⠂ ⠂ ⠒ ⠚ ⠙ ⠉ ⠁  |
| dots7 | 80 | ⠈ ⠉ ⠋ ⠓ ⠒ ⠐ ⠐ ⠒ ⠖ ⠦ ⠤ ⠠ ⠠ ⠤ ⠦ ⠖ ⠒ ⠐ ⠐ ⠒ ⠓ ⠋ ⠉ ⠈  |
| dots8 | 80 | ⠁ ⠁ ⠉ ⠙ ⠚ ⠒ ⠂ ⠂ ⠒ ⠲ ⠴ ⠤ ⠄ ⠄ ⠤ ⠠ ⠠ ⠤ ⠦ ⠖ ⠒ ⠐ ⠐ ⠒ ⠓ ⠋ ⠉ ⠈ ⠈  |
| dots9 | 80 | ⢹ ⢺ ⢼ ⣸ ⣇ ⡧ ⡗ ⡏  |
| dqpb | 100 | d q p b  |
| earth | 180 | 🌍  🌎  🌏   |
| flip | 70 | _ _ _ - ` ` &#39; ´ - _ _ _  |
| growHorizontal | 120 | ▏ ▎ ▍ ▌ ▋ ▊ ▉ ▊ ▋ ▌ ▍ ▎  |
| growVertical | 120 | ▁ ▃ ▄ ▅ ▆ ▇ ▆ ▅ ▄ ▃  |
| hamburger | 100 | ☱ ☲ ☴  |
| hearts | 100 | 💛  💙  💜  💚  ❤️   |
| line | 130 | - \ | /  |
| line2 | 100 | ⠂ - – — – -  |
| monkey | 300 | 🙈  🙈  🙉  🙊   |
| moon | 80 | 🌑  🌒  🌓  🌔  🌕  🌖  🌗  🌘   |
| noise | 100 | ▓ ▒ ░  |
| pipe | 100 | ┤ ┘ ┴ └ ├ ┌ ┬ ┐  |
| pong | 80 | ▐⠂       ▌ ▐⠈       ▌ ▐ ⠂      ▌ ▐ ⠠      ▌ ▐  ⡀     ▌ ▐  ⠠     ▌ ▐   ⠂    ▌ ▐   ⠈    ▌ ▐    ⠂   ▌ ▐    ⠠   ▌ ▐     ⡀  ▌ ▐     ⠠  ▌ ▐      ⠂ ▌ ▐      ⠈ ▌ ▐       ⠂▌ ▐       ⠠▌ ▐       ⡀▌ ▐      ⠠ ▌ ▐      ⠂ ▌ ▐     ⠈  ▌ ▐     ⠂  ▌ ▐    ⠠   ▌ ▐    ⡀   ▌ ▐   ⠠    ▌ ▐   ⠂    ▌ ▐  ⠈     ▌ ▐  ⠂     ▌ ▐ ⠠      ▌ ▐ ⡀      ▌ ▐⠠       ▌  |
| runner | 140 | 🚶  🏃   |
| shark | 120 | ▐|\____________▌ ▐_|\___________▌ ▐__|\__________▌ ▐___|\_________▌ ▐____|\________▌ ▐_____|\_______▌ ▐______|\______▌ ▐_______|\_____▌ ▐________|\____▌ ▐_________|\___▌ ▐__________|\__▌ ▐___________|\_▌ ▐____________|\▌ ▐____________/|▌ ▐___________/|_▌ ▐__________/|__▌ ▐_________/|___▌ ▐________/|____▌ ▐_______/|_____▌ ▐______/|______▌ ▐_____/|_______▌ ▐____/|________▌ ▐___/|_________▌ ▐__/|__________▌ ▐_/|___________▌ ▐/|____________▌  |
| simpleDots | 400 | .   ..  ...      |
| simpleDotsScrolling | 200 | .   ..  ...  ..   .      |
| smiley | 200 | 😄  😝   |
| squareCorners | 180 | ◰ ◳ ◲ ◱  |
| squish | 100 | ╫ ╪  |
| star | 70 | ✶ ✸ ✹ ✺ ✹ ✷  |
| star2 | 80 | &#43; x *  |
| toggle | 250 | ⊶ ⊷  |
| toggle10 | 100 | ㊂ ㊀ ㊁  |
| toggle11 | 50 | ⧇ ⧆  |
| toggle12 | 120 | ☗ ☖  |
| toggle13 | 80 | = * -  |
| toggle2 | 80 | ▫ ▪  |
| toggle3 | 120 | □ ■  |
| toggle4 | 100 | ■ □ ▪ ▫  |
| toggle5 | 100 | ▮ ▯  |
| toggle6 | 300 | ဝ ၀  |
| toggle7 | 80 | ⦾ ⦿  |
| toggle8 | 100 | ◍ ◌  |
| toggle9 | 100 | ◉ ◎  |
| triangle | 50 | ◢ ◣ ◤ ◥  |
| weather | 100 | ☀️  ☀️  ☀️  🌤  ⛅️  🌥  ☁️  🌧  🌨  🌧  🌨  🌧  🌨  ⛈  🌨  🌧  🌨  ☁️  🌥  ⛅️  🌤  ☀️  ☀️   |