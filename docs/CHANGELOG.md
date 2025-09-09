| Release | Date       | Comments                                                                                                                          |
|---------|------------|-----------------------------------------------------------------------------------------------------------------------------------|
| 2.3.3   | 2025.08.25 | Removed extra newline + `:` at the end of the Error() output                                                                      |
| 2.3.2   | 2025.08.14 | Fixed issue where Fatal() is not properly returning its string value                                                              |                                                            
| 2.3.1   | 2025.08.14 | New helper wrapper, ErrorNoColor() to handle NoColour as simply as possible from the caller software<br>GO version bump to 1.25.0 |
| 2.3.0   | 2025.08.14 | now handling colourized output, or none (for logging facilities)                                                                  |
| 2.2.0   | 2025.08.09 | added a string method to "stringify" the error type<br>removed the logging subpackage<br>GO version bump                          |
| 2.1.1   | 2025.07.23 | interface handling is not working as expected, found a workaround                                                                 |
| 2.1.0   | 2025.07.23 | Log() now accepts an interface so we can send anything we wish to the logger                                                      |
| 2.0.0   | 2025.07.20 | added a logging subpackage                                                                                                        |
| 1.7.0   | 2024.05.05 | fatal error is now the default error type                                                                                         |
| 1.6.7   | 2024.05.02 | type assertion failures are handled within the package                                                                            |
| 1.6.6   | 2024.04.29 | a fatal error will now halt the program, as it should                                                                             |
| 1.6.5   | 2024.04.03 | getting ready for param initialization with a (so far, empty) constructor                                                         |
| 1.5.1   | 2024.04.03 | fixed gh action                                                                                                                   |
| 1.5.0   | 2024.04.02 | fixed output issues                                                                                                               |
| 1.1.0   | 2024.04.01 | corrections to output<br> Added main.go for debugging purposes<br>(this file is wholly commented in prod)                         |
| 1.0.1   | 2024.04.01 | fixed minor typo.                                                                                                                 |
| 1.0.0   | 2024.03.31 | initial version.                                                                                                                  |




