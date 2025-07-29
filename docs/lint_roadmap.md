# Lint Roadmap

This document outlines various linting issues detected in the codebase. Each category of issues has been grouped accordingly.

| Lint Type   | File                                | Line | Issue Description                                                                                                           |
|-------------|-------------------------------------|------|-----------------------------------------------------------------------------------------------------------------------------|
| **revive**  |
|             | auth.go                             | 13   | exported: comment on exported method Client.AuthPartnerLogin should be of the form "AuthPartnerLogin ..."                  |
|             | auth.go                             | 55   | exported: comment on exported method Client.AuthUserLogin should be of the form "AuthUserLogin ..."                        |
|             | gopiano.go                          | 12   | package-comments: package comment is detached; there should be no blank lines between it and the package statement          |
|             | gopiano.go                          | 30   | exported: comment on exported type ClientDescription should be of the form "ClientDescription ..."                         |
|             | gopiano.go                          | 41   | exported: comment on exported var AndroidClient should be of the form "AndroidClient ..."                                  |
|             | gopiano.go                          | 42   | var-declaration: should omit type ClientDescription from declaration of var AndroidClient; it will be inferred from the RHS |
|             | gopiano.go                          | 52   | exported: comment on exported type Client should be of the form "Client ..."                                               |
|             | gopiano.go                          | 65   | exported: comment on exported function NewClient should be of the form "NewClient ..."                                     |
|             | gopiano.go                          | 119  | exported: comment on exported method Client.PandoraCall should be of the form "PandoraCall ..."                            |
|             | gopiano.go                          | 178  | exported: comment on exported method Client.BlowfishCall should be of the form "BlowfishCall ..."                           |
|             | gopiano.go                          | 189  | exported: comment on exported method Client.GetSyncTime should be of the form "GetSyncTime ..."                             |
|             | misc.go                             | 11   | exported: comment on exported method Client.ExplainTrack should be of the form "ExplainTrack ..."                           |
|             | misc.go                             | 34   | exported: comment on exported method Client.MusicSearch should be of the form "MusicSearch ..."                             |
|             | misc.go                             | 56   | exported: comment on exported method Client.BookmarkAddArtistBookmark should be of the form "BookmarkAddArtistBookmark ..."|
|             | misc.go                             | 79   | exported: comment on exported method Client.BookmarkAddSongBookmark should be of the form "BookmarkAddSongBookmark ..."    |
|             | requests\requests.go                | 1    | package-comments: package comment should be of the form "Package requests ..."                                             |
|             | requests\requests.go                | 6    | exported: exported type AuthPartnerLogin should have comment or be unexported                                               |
|             | requests\requests.go                | 14   | exported: exported type AuthUserLogin should have comment or be unexported                                                   |
|             | requests\requests.go                | 35   | exported: exported type UserGetBookmarks should have comment or be unexported                                                |
|             | requests\requests.go                | 36   | exported: exported type UserGetStationListChecksum should have comment or be unexported                                       |
|             | requests\requests.go                | 37   | exported: exported type UserCanSubscribe should have comment or be unexported                                                 |
|             | requests\requests.go                | 40   | exported: exported type UserCreateUser should have comment or be unexported                                                   |
|             | requests\requests.go                | 54   | exported: exported type UserEmailPassword should have comment or be unexported                                                |
|             | requests\requests.go                | 60   | exported: exported type UserGetStationList should have comment or be unexported                                               |
|             | requests\requests.go                | 66   | exported: exported type UserSetQuickMix should have comment or be unexported                                                  |
|             | requests\requests.go                | 78   | exported: exported type UserSleepSong should have comment or be unexported                                                    |
|             | requests\requests.go                | 79   | exported: exported type BookmarkAddArtistBookmark should have comment or be unexported                                        |
|             | requests\requests.go                | 80   | exported: exported type BookmarkAddSongBookmark should have comment or be unexported                                          |
|             | requests\requests.go                | 83   | exported: exported type MusicSearch should have comment or be unexported                                                       |
|             | requests\requests.go                | 89   | exported: exported type StationCreateStation should have comment or be unexported                                             |
|             | requests\requests.go                | 97   | exported: exported type StationDeleteStation should have comment or be unexported                                              |
|             | requests\requests.go                | 103  | exported: exported type StationAddFeedback should have comment or be unexported                                                |
|             | requests\requests.go                | 110  | exported: exported type StationDeleteFeedback should have comment or be unexported                                              |
|             | requests\requests.go                | 116  | exported: exported type StationAddMusic should have comment or be unexported                                                    |
|             | requests\requests.go                | 123  | exported: exported type StationDeleteMusic should have comment or be unexported                                                 |
|             | requests\requests.go                | 130  | exported: exported type StationGetGenreStations should have comment or be unexported                                            |
|             | requests\requests.go                | 131  | exported: exported type StationGetGenreStationsChecksum should have comment or be unexported                                    |
|             | requests\requests.go                | 134  | exported: exported type StationGetPlaylist should have comment or be unexported                                                 |
|             | requests\requests.go                | 140  | exported: exported type StationGetStation should have comment or be unexported                                                  |
|             | requests\requests.go                | 147  | exported: exported type StationShareStation should have comment or be unexported                                                |
|             | requests\requests.go                | 155  | exported: exported type StationRenameStation should have comment or be unexported                                               |
|             | requests\requests.go                | 162  | exported: exported type StationTransformSharedStation should have comment or be unexported                                       |
|             | requests\requests.go                | 168  | exported: exported type ExplainTrack should have comment or be unexported                                                       |
|             | responses\responses.go              | 1    | package-comments: package comment should be of the form "Package responses ..."                                              |
|             | responses\responses.go              | 11   | var-declaration: should omit type map[int]string from declaration of var ErrorCodeMap; it will be inferred from the RHS        |
|             | responses\responses.go              | 58   | exported: exported type ErrorResponse should have comment or be unexported                                                     |
|             | responses\responses.go              | 85   | exported: comment on exported method DateResponse.GetDate should be of the form "GetDate ..."                                 |
|             | responses\responses.go              | 91   | exported: exported type AuthPartnerLogin should have comment or be unexported                                                  |
|             | responses\responses.go              | 110  | exported: exported type AuthUserLogin should have comment or be unexported                                                      |
|             | responses\responses.go              | 115  | var-naming: struct field ListeningTimeoutAlertMsgUri should be ListeningTimeoutAlertMsgURI                                     |
|             | responses\responses.go              | 130  | exported: exported type UserCanSubscribe should have comment or be unexported                                                   |
|             | responses\responses.go              | 137  | exported: exported type UserCreateUser should have comment or be unexported                                                     |
|             | responses\responses.go              | 139  | exported: exported type ArtistBookmark should have comment or be unexported                                                     |
|             | responses\responses.go              | 147  | exported: exported type BookmarkAddArtistBookmark should have comment or be unexported                                          |
|             | responses\responses.go              | 151  | exported: exported type SongBookmark should have comment or be unexported                                                       |
|             | responses\responses.go              | 163  | exported: exported type BookmarkAddSongBookmark should have comment or be unexported                                             |
|             | responses\responses.go              | 167  | exported: exported type Station should have comment or be unexported                                                            |
|             | responses\responses.go              | 203  | exported: exported type StationList should have comment or be unexported                                                        |
|             | responses\responses.go              | 218  | exported: exported type UserGetBookmarks should have comment or be unexported                                                   |
|             | responses\responses.go              | 225  | exported: exported type UserGetStationList should have comment or be unexported                                                  |
|             | responses\responses.go              | 232  | exported: exported type UserGetStationListChecksum should have comment or be unexported                                          |
|             | responses\responses.go              | 238  | exported: exported type MusicSearch should have comment or be unexported                                                        |
|             | responses\responses.go              | 257  | exported: exported type FeedbackResponse should have comment or be unexported                                                    |
|             | responses\responses.go              | 265  | exported: exported type StationAddFeedback should have comment or be unexported                                                  |
|             | responses\responses.go              | 269  | exported: exported type StationAddMusic should have comment or be unexported                                                     |
|             | responses\responses.go              | 277  | exported: exported type StationResponse should have comment or be unexported                                                     |
|             | responses\responses.go              | 281  | exported: exported type StationCreateStation should have comment or be unexported                                                |
|             | responses\responses.go              | 282  | exported: exported type StationGetStation should have comment or be unexported                                                   |
|             | responses\responses.go              | 283  | exported: exported type StationRenameStation should have comment or be unexported                                                |
|             | responses\responses.go              | 284  | exported: exported type StationTransformSharedStation should have comment or be unexported                                        |
|             | responses\responses.go              | 287  | exported: exported type StationGetGenreStations should have comment or be unexported                                             |
|             | responses\responses.go              | 300  | exported: exported type StationGetGenreStationsChecksum should have comment or be unexported                                     |
|             | responses\responses.go              | 306  | exported: exported type StationGetPlaylist should have comment or be unexported                                                  |
|             | responses\responses.go              | 341  | exported: exported type ExplainTrack should have comment or be unexported                                                        |
|             | station.go                          | 11   | exported: comment on exported method Client.StationAddFeedback should be of the form "StationAddFeedback ..."                 |
|             | station.go                          | 36   | exported: comment on exported method Client.StationAddMusic should be of the form "StationAddMusic ..."                        |
|             | station.go                          | 61   | exported: comment on exported method Client.StationCreateStationTrack should be of the form "StationCreateStationTrack ..."   |
|             | station.go                          | 86   | exported: comment on exported method Client.StationCreateStationMusic should be of the form "StationCreateStationMusic ..."   |
|             | station.go                          | 109  | exported: comment on exported method Client.StationDeleteFeedback should be of the form "StationDeleteFeedback ..."           |
|             | station.go                          | 126  | exported: comment on exported method Client.StationDeleteMusic should be of the form "StationDeleteMusic ..."                 |
|             | station.go                          | 143  | exported: comment on exported method Client.StationDeleteStation should be of the form "StationDeleteStation ..."             |
|             | station.go                          | 160  | exported: comment on exported method Client.StationGetGenreStations should be of the form "StationGetGenreStations ..."       |
|             | station.go                          | 181  | exported: comment on exported method Client.StationGetPlaylist should be of the form "StationGetPlaylist ..."                 |
|             | station.go                          | 205  | exported: comment on exported method Client.StationGetStation should be of the form "StationGetStation ..."                   |
|             | station.go                          | 230  | exported: comment on exported method Client.StationShareStation should be of the form "StationShareStation ..."               |
|             | station.go                          | 252  | exported: comment on exported method Client.StationRenameStation should be of the form "StationRenameStation ..."             |
|             | station.go                          | 275  | exported: comment on exported method Client.StationTransformSharedStation should be of the form "StationTransformSharedStation ..." |
|             | user.go                             | 11   | exported: comment on exported method Client.UserCanSubscribe should be of the form "UserCanSubscribe ..."                     |
|             | user.go                             | 33   | exported: comment on exported method Client.UserCreateUser should be of the form "UserCreateUser ..."                         |
|             | user.go                             | 69   | exported: comment on exported method Client.UserEmailPassword should be of the form "UserEmailPassword ..."                   |
|             | user.go                             | 86   | exported: comment on exported method Client.UserGetBookmarks should be of the form "UserGetBookmarks ..."                     |
|             | user.go                             | 108  | exported: comment on exported method Client.UserGetStationList should be of the form "UserGetStationList ..."                 |
|             | user.go                             | 131  | exported: comment on exported method Client.UserGetStationListChecksum should be of the form "UserGetStationListChecksum ..." |
|             | user.go                             | 153  | exported: comment on exported method Client.UserSetQuickMix should be of the form "UserSetQuickMix ..."                       |
|             | user.go                             | 170  | exported: comment on exported method Client.UserSleepSong should be of the form "UserSleepSong ..."                           |
| **errcheck**|
|             | encrypt_test.go                     | 6    | Error return value is not checked                                                                                           |
|             | encrypt_test.go                     | 18   | Error return value is not checked                                                                                           |
| **bodyclose**|
|             | gopiano.go                          | 149  | response body must be closed                                                                                                |
| **staticcheck**|
|             | gopiano.go                          | 19   | SA1019: "io/ioutil" has been deprecated since Go 1.19: As of Go 1.16, the same functionality is now provided by package [io] or package [os], and those implementations should be preferred in new code. See the specific function documentation for details. |
|             | gopiano.go                          | 25   | SA1019: "golang.org/x/crypto/blowfish" is deprecated: any new system should use AES or XChaCha20-Poly1305.                  |
| **style/naming/other**|
|             | responses\responses.go              | 88   | Magic number: 60, in argument detected                                                                            |
|             | gopiano.go                          | 142  | net/http.NewRequest must not be called. use net/http.NewRequestWithContext                                                   |
|             | gopiano.go                          | 140  | var-naming: var callUrl should be callURL                                                                                    |
|             | responses\responses.go              | 115  | var-naming: struct field ListeningTimeoutAlertMsgUri should be ListeningTimeoutAlertMsgURI                                   |
|             | requests\requests.go                | 21   | json(camel): got 'IncludeDemographics' want 'includeDemographics'                                                            |
|             | auth.go                             | 41   | unnecessary conversion                                                                                                                                            |

