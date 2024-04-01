# customError
___
A simple extension on the Error() interface. That provides output in yellow (warning) or red (error), according to the error type.

# *__IMPORTANT NOTE__*
___
This github repo is aread-only copy of the actual *source of truth* repo. The source repo is outside of github, and is the one where all Pull Requests and merges happen.

If you need changes, please fill an issue in the issues tab of the repo with what needs to be corrected, or better yet, include a patch to whichever files need changes.

My main repo is on gitea (hosted somewhere else), and compability issues between gitea and github prevent me to `git pull` the github repo into the gitea repo, but I can push from gitea to github, thanks to git remote add *githubRepo ghRepoURL*

In a good day, in a good mood, I __could__ give access to some users to my gitea repo, but... well.. most likely not. Accept it, or not: by no means this package is a perfect package, we're all humans with different levels of understanding of GO or whatnot, so it does not suit you ? Please do find another more suitable solution to your needs.
## Structure

We have this very simple (... for now) structure and its dependencies. I expect to expand this type down the road.
```go
type ErrortypeIota int

const (
	Undefined ErrortypeIota = iota
	Fatal
	Warning
	Continuable
)

type CustomError struct {
    Title    string         // optional
    Message  string
    Code     int           // optional
    Fatality ErrortypeIota // if omitted, we will use "Undefined, and will throw a panic() right away, so... define it :)
}
```

If you wish to customize an error message that would otherwise be plain, or just _meh_, you just need to populate a variable of the above struct type with the proper members, and use $YOUR_VAR_NAME.Error().
More on that later on.


### Struct members

- Of the above structure, what is important here is the `Fatality` member from the CustomError struct. It is from there that we decide of the colors used in the output.

When this member is not explicitely initialized, it assumes the value of "Undefined", and will throw, eventually a panic() and stops the software. Otherwise, this is the value that determines which colors are used to display the error message.

- The `Title` struct member is self-explanatory. It prepends the actual error message with a title, such as (example): "FATAL: ", or "WARNING", or whatever.

The color of `Title` is assumed from `Fatality`, which is red. If omitted, the actual error message (member `Message`) will be of the proper color.

- The `Message` member is the "real" error message. A safe bet is to have it using `err.Error()`
Actually, if you omit to define this member, it will be set to "Unspecified error" as `Message`.


- `Code` is totally optional, and is there for totally real customized error messages where you define your own errors.

It will not show up in the output if omitted.


## Usage

Assume you have the following code that returns a plain error:

```go
func myFunc() error {
    if err := os.MkDirAll(var1, var2); err != nil {
        return err
    }
}
```

This would return an error the usual way to whichever function called `myFunc()`.

With the CustomError package, it gets a bit more colourful, and explicit:

1. You first need to download the package:
`go get "github.com/jeanfrancoisgratton/customErrors"`
2. In the import section of your file, you need to add the above package in that section
3. Then, you will need to change your previous code a bit:

```go
func myFunc() error {
    if err := os.MkDirAll(var1, var2); err != nil {
        ce := CustomError {Fatality: ErrortypeIota.Fatal, Message: err.Error(), Title: "Error creating directory", Code: -8}
        return ce.Error()		
    }
}
```

__Remember__ :

You can omit the `Title`, `Message` and `Code` members of the struct without any hitch;
- if you omit `Fatality`, your software will exit with a panic() call.
- if you omit `Message`, it will use the error message from `err.Error()` (as it would have done before using CustomError)
- if you omit `Title`, the coloured output will display `Message`, without any title
- if you omit `Code`, it just gets discarded from the coloured output

