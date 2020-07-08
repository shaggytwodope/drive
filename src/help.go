// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drive

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	prettywords "github.com/odeke-em/pretty-words"
)

const (
	AboutKey                  = "about"
	AllKey                    = "all"
	CopyKey                   = "copy"
	DeleteKey                 = "delete"
	DeInitKey                 = "deinit"
	EditDescriptionKey        = "edit-description"
	EditDescriptionShortKey   = "edit-desc"
	ServiceAccountJSONFileKey = "service-account-file"
	DiffKey                   = "diff"
	AddressKey                = "address"
	EmptyTrashKey             = "emptytrash"
	FeaturesKey               = "features"
	HelpKey                   = "help"
	InitKey                   = "init"
	LinkKey                   = "Link"
	ListKey                   = "list"
	DuKey                     = "du"
	Md5sumKey                 = "md5sum"
	MoveKey                   = "move"
	OcrKey                    = "ocr"
	ConvertKey                = "convert"
	OSLinuxKey                = "linux"
	PullKey                   = "pull"
	PipedKey                  = "piped"
	PushKey                   = "push"
	PubKey                    = "pub"
	QRLinkKey                 = "qr"
	RenameKey                 = "rename"
	QuotaKey                  = "quota"
	ShareKey                  = "share"
	StatKey                   = "stat"
	TouchKey                  = "touch"
	TrashKey                  = "trash"
	UnshareKey                = "unshare"
	UntrashKey                = "untrash"
	UnpubKey                  = "unpub"
	VersionKey                = "version"
	NewKey                    = "new"
	IndexKey                  = "index"
	PruneKey                  = "prune"
	StarKey                   = "star"
	UnStarKey                 = "unstar"

	CoercedMimeKeyKey        = "coerced-mime"
	ExportsKey               = "export"
	ExportsDirKey            = "exports-dir"
	NoClobberKey             = "no-clobber"
	RecursiveKey             = "recursive"
	IgnoreChecksumKey        = "ignore-checksum"
	ExcludeOpsKey            = "exclude-ops"
	IgnoreConflictKey        = "ignore-conflict"
	IgnoreNameClashesKey     = "ignore-name-clashes"
	ClashesKey               = "clashes"
	CommentStr               = "#"
	DepthKey                 = "depth"
	EmailsKey                = "emails"
	EmailMessageKey          = "emailMessage"
	ForceKey                 = "force"
	QuietKey                 = "quiet"
	IdKey                    = "id"
	QuitShortKey             = "q"
	YesShortKey              = "Y"
	QuitLongKey              = "quit"
	MatchesKey               = "matches"
	HiddenKey                = "hidden"
	Md5Key                   = "md5"
	NoPromptKey              = "no-prompt"
	SizeKey                  = "size"
	NameKey                  = "name"
	OpenKey                  = "open"
	OriginalNameKey          = "oname"
	ModTimeKey               = "modt"
	LastViewedByMeTimeKey    = "lvt"
	AccountTypeKey           = "account-type"
	RoleKey                  = "role"
	TypeKey                  = "type"
	TrashedKey               = "trashed"
	SkipMimeKeyKey           = "skip-mime"
	MatchMimeKeyKey          = "exact-mime"
	ExactTitleKey            = "exact-title"
	MatchOwnerKey            = "match-owner"
	ExactOwnerKey            = "exact-owner"
	NotOwnerKey              = "skip-owner"
	SortKey                  = "sort"
	FolderKey                = "folder"
	MimeKey                  = "mime-key"
	PageSizeKey              = "pagesize"
	DriveRepoRelPath         = "github.com/odeke-em/drive"
	UrlKey                   = "url"
	ReportIssueKey           = "report-issue"
	IssueTitleKey            = "title"
	IssueBodyKey             = "body"
	SkipContentCheckKey      = "skip-content-check"
	TouchModTimeKey          = "time"
	TouchTimeFmtSpecifierKey = "format"
	TouchOffsetDurationKey   = "duration"
)

const (
	DescAbout                 = "print out information about your Google drive"
	DescAll                   = "print out the entire help section"
	DescAllStarred            = "all the starred files"
	DescCopy                  = "copy remote paths to a destination"
	DescDelete                = "deletes the items permanently. This operation is irreversible"
	DescDiff                  = "compares local files with their remote equivalent"
	DescEdit                  = "edit the attributes of a file"
	DescEmptyTrash            = "permanently cleans out your trash"
	DescExcludeOps            = "exclude operations"
	DescFeatures              = "returns information about the features of your drive"
	DescIndex                 = "fetch indices from remote"
	DescHelp                  = "Get help for a topic"
	DescInit                  = "initializes a directory and authenticates user"
	DescDeInit                = "removes the user's credentials and initialized files"
	DescList                  = "lists the contents of remote path"
	DescMove                  = "move files/folders"
	DescPiped                 = "get content in from standard input (stdin)"
	DescQuota                 = "prints out information related to your quota space"
	DescPublish               = "publishes a file and prints its publicly available url"
	DescRename                = "renames a file/folder"
	DescPull                  = "pulls remote changes from Google Drive"
	DescPruneIndices          = "remove stale indices"
	DescPush                  = "push local changes to Google Drive"
	DescShare                 = "share files with specific emails giving the specified users specifies roles and permissions"
	DescStar                  = "star files"
	DescUnStar                = "unstar files"
	DescStat                  = "display information about a file"
	DescTouch                 = "updates a remote file's modification time to that currently on the server"
	DescTrash                 = "moves files to trash"
	DescUnshare               = "revoke a user's access to a file"
	DescUntrash               = "restores files from trash to their original locations"
	DescUnpublish             = "revokes public access to a file"
	DescVersion               = "prints the version"
	DescMd5sum                = "prints a list compatible with md5sum(1)"
	DescDu                    = "similar to util `du` gives you disk usage"
	DescAccountTypes          = "\n\t* anyone.\n\t* user.\n\t* domain.\n\t* group"
	DescRoles                 = "\n\t* owner.\n\t* reader.\n\t* writer.\n\t* commenter."
	DescExplicitylPullExports = "explicitly pull exports"
	DescIgnoreChecksum        = "avoids computation of checksums as a final check." +
		"\nUse cases may include:\n\t* when you are low on bandwidth e.g SSHFS." +
		"\n\t* Are on a low power device"
	DescIgnoreConflict               = "turns off the conflict resolution safety"
	DescIgnoreNameClashes            = "ignore name clashes"
	DescSort                         = "sort items by a combination of attributes\n\t* modtime.\n\t* md5.\n\t* name.\n\t* size.\n\t* type.\n\t* version\ncomma separated e.g modtime,md5_r,name"
	DescSkipMime                     = "skip elements with mimeTypes derived from these extensions"
	DescMatchMime                    = "get elements with the exact mimeTypes derived from extensions"
	DescMatchTitle                   = "elements with matching titles"
	DescExactTitle                   = "get elements with the exact titles"
	DescMatchOwner                   = "elements with matching owners"
	DescExactOwner                   = "elements with the exact owner"
	DescNotOwner                     = "ignore elements owned by these users"
	DescNew                          = "create a new file/folder"
	DescAllIndexOperations           = "perform all the index related operations"
	DescOpen                         = "open a file in the appropriate filemanager or default browser"
	DescUrl                          = "returns the remote URL of each file"
	DescVerbose                      = "show step by step information verbosely"
	DescFixClashes                   = "fix clashes by renaming or trashing all files"
	DescFixClashesMode               = "set fix policy to rename or trash"
	DescListClashes                  = "list clashes"
	DescDescription                  = "set the description"
	DescQR                           = "open up the QR code for specified files"
	DescStarred                      = "operate only on starred files"
	DescUnifiedDiff                  = "unified diff"
	DescDiffBaseLocal                = "when set uses local as the base other remote will be used as the base"
	DescClashesOpById                = "operate on clashes by id instead of by path"
	DescIssueBody                    = "the detailed description of the issue being filed"
	DescIssueTitle                   = "the title of the issue being filed"
	DescReportIssue                  = "report an issue to the project's issue tracker"
	DescId                           = "retrieve the fileId for the specified paths"
	DescSkipContentCheck             = "skip diffing actual body content, show only name, time, type changes"
	DescPushDestination              = "specify the final destination of the contents of an operation"
	DescExponentialBackoffRetryCount = "max number of retries for exponential backoff"
	DescEncryptionPassword           = "encryption password"
	DescDecryptionPassword           = "decryption password"
	DescWithLink                     = "turn off file indexing so that only those with the link can view it"
	DescAllowDesktopLinks            = "allows docs + sheets to be pulled as .desktop files or URL linked files"
	DescKeepParent                   = "ensures that when moving a file into a destination, that we also retain its original parent so that it will exist in more than one folder"

	DescTouchTimeStr          = "the time each file's modification time should be set to"
	DescTouchOffsetDuration   = "the duration offset from now that each file's modification time should be set to e.g -32h\nSee https://golang.org/pkg/time/#ParseDuration"
	DescTouchTimeFmtSpecifier = "the custom layout that you'd like your time to be set in, representative of the way 'Mon Jan 2 15:04:05 -0700 MST 2006' should be represented\nSee https://golang.org/pkg/time/#Parse"
)

const (
	CLIOptionDescription        = "description"
	CLIOptionExplicitlyExport   = "explicitly-export"
	CLIOptionIgnoreChecksum     = "ignore-checksum"
	CLIOptionIgnoreConflict     = "ignore-conflict"
	CLIOptionIgnoreNameClashes  = "ignore-name-clashes"
	CLIOptionExcludeOperations  = "exclude-ops"
	CLIOptionId                 = "id"
	CLIOptionNoClobber          = "no-clobber"
	CLIOptionNotify             = "notify"
	CLIOptionSkipMime           = "skip-mime"
	CLIOptionExactMime          = "exact-mime"
	CLIOptionMatchMime          = "match-mime"
	CLIOptionExactTitle         = "exact-title"
	CLIOptionMatchTitle         = "match-title"
	CLIOptionExactOwner         = "exact-owner"
	CLIOptionMatchOwner         = "match-owner"
	CLIOptionNotOwner           = "skip-owner"
	CLIOptionPruneIndices       = "prune"
	CLIOptionAllIndexOperations = "all-ops"
	CLIOptionVerboseKey         = "verbose"
	CLIOptionVerboseShortKey    = "v"
	CLIOptionOpen               = "open"
	CLIOptionWebBrowser         = "web-browser"
	CLIOptionFileBrowser        = "file-browser"
	CLIOptionDirectories        = "directories"
	CLIOptionFiles              = "files"
	CLIOptionLongFmt            = "long"
	CLIOptionFixClashesKey      = "fix-clashes"
	CLIOptionPiped              = "piped"
	CLIOptionStarred            = "starred"
	CLIOptionAllStarred         = "all"
	CLIOptionUnified            = "unified"
	CLIOptionUnifiedShortKey    = "u"
	CLIOptionDiffBaseLocal      = "base-local"
	CLIOptionFixClashes         = "fix"
	CLIOptionFixClashesMode     = "fix-mode"
	CLIOptionListClashes        = "list"
	CLIOptionPushDestination    = "destination"
	CLIOptionRenameLocal        = "local"
	CLIOptionRenameRemote       = "remote"
	CLIOptionRetryCount         = "retry-count"
	CLIEncryptionPassword       = "encryption-password"
	CLIDecryptionPassword       = "decryption-password"
	CLIOptionWithLink           = "with-link"
	CLIOptionDesktopLinks       = "desktop-links"
	CLIOptionKeepParent         = "keep-parent"

	CLIOptionUploadChunkSize = "upload-chunk-size"
	CLIOptionUploadRateLimit = "upload-rate-limit"

	CLIOptionExportsDumpToSameDirectory = "same-exports-dir"

	CLIOptionTrashed = TrashedKey
)

const (
	DefaultMaxTraversalDepth = -1
)

const (
	GoogleApiClientIdEnvKey     = "GOOGLE_API_CLIENT_ID"
	GoogleApiClientSecretEnvKey = "GOOGLE_API_CLIENT_SECRET"
	DriveGoMaxProcsKey          = "DRIVE_GOMAXPROCS"
	GoMaxProcsKey               = "GOMAXPROCS"
)

const (
	DesktopExtension = "desktop"
)

const (
	InfiniteDepth = -1
)

var (
	skipChecksumNote = fmt.Sprintf(
		"\nNote: You can skip checksum verification by passing in flag `-%s`", CLIOptionIgnoreChecksum)
	PermanentDeletionNoPromptError = fmt.Errorf("%q is set yet performing a permanent deletion. Please see issue https://github.com/odeke-em/drive/issues/448", NoPromptKey)
)

var docMap = map[string][]string{
	AboutKey: []string{
		DescAbout,
	},
	CopyKey: []string{
		DescCopy,
	},
	DeleteKey: []string{
		DescDelete,
	},
	DiffKey: []string{
		DescDiff, "Accepts multiple remote paths for line by line comparison",
		skipChecksumNote,
	},
	EditDescriptionShortKey: []string{
		DescEdit, "Accepts multiple remote paths as well as ids",
	},
	EmptyTrashKey: []string{
		DescEmptyTrash,
	},
	FeaturesKey: []string{
		DescFeatures,
	},
	InitKey: []string{
		DescInit, "Requests for access to your Google Drive",
		"Creating a folder that contains your credentials",
		"Note: `init` in an already initialized drive will erase the old credentials",
	},
	PullKey: []string{
		DescPull, "Downloads content from the remote drive or modifies",
		" local content to match that on your Google Drive",
		skipChecksumNote,
	},
	PushKey: []string{
		DescPush, "Uploads content to your Google Drive from your local path",
		"Push comes in a couple of flavors",
		"\t* Ordinary push: `drive push path1 path2 path3`",
		"\t* Mounted push: `drive push -m path1 [path2 path3] drive_context_path`",
		skipChecksumNote,
	},
	ListKey: []string{
		DescList,
		"List the information of a remote path not necessarily present locally",
		"Allows printing of long options and by default does minimal printing",
	},
	MoveKey: []string{
		DescMove,
		"Moves files/folders between folders",
	},
	OpenKey: []string{
		DescOpen, fmt.Sprintf("toggle between %q=bool and %q=bool",
			CLIOptionWebBrowser, CLIOptionFileBrowser),
	},
	PubKey: []string{
		DescPublish, "Accepts multiple paths",
	},
	RenameKey: []string{
		DescRename, "Accepts <src> <newName>",
	},
	QuotaKey: []string{DescQuota},
	ShareKey: []string{
		DescShare, "Accepts multiple paths",
		"Specify the emails to share with as well as the message to send them on notification",
		"Accepted values for:\n+ accountType: ",
		DescAccountTypes, "\n+ roles:", DescRoles,
	},
	StatKey: []string{
		DescStat, "provides detailed information about a remote file",
		"Accepts multiple paths",
	},
	TouchKey: []string{
		DescTouch, "Given a list of remote files `touch` updates their",
		"last edit times to that currently on the server",
	},
	TrashKey: []string{
		DescTrash, "Sends a list of remote files to trash",
	},
	UnshareKey: []string{
		DescUnshare, "Accepts multiple paths",
		"Accepted values for accountTypes::", DescAccountTypes,
	},
	UntrashKey: []string{
		DescUntrash, "takes remote files out of the trash",
		"Note: untrash is a relative path command so any resolutions are made",
		"relative to the current working directory i.e",
		"\n\t$ drive trash mnt/logos",
	},
	UnpubKey: []string{
		DescUnpublish, "revokes public access to a list of remote files",
	},
	UrlKey: []string{
		DescUrl, "takes multiple paths or ids",
	},
	VersionKey: []string{
		DescVersion, fmt.Sprintf("current version is: %s", Version),
	},
}

func createAndRegisterAliases() map[string][]string {
	aliases := map[string][]string{
		CopyKey:            []string{"cp"},
		ListKey:            []string{"ls"},
		MoveKey:            []string{"mv"},
		DeleteKey:          []string{"del"},
		EditDescriptionKey: []string{EditDescriptionShortKey},
		IdKey:              []string{"file-id"},
		ReportIssueKey:     []string{"issue", "report"},
	}

	for originalKey, aliasList := range aliases {
		docDetails, ok := docMap[originalKey]
		if !ok {
			continue
		}

		for _, alias := range aliasList {
			docMap[alias] = docDetails
		}
	}

	return aliases
}

var Aliases = createAndRegisterAliases()

func ShowAllDescriptions() {
	keys := []string{}
	for key, _ := range docMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		ShowDescription(key)
	}
}

func ShowDescriptions(topics ...string) {
	if len(topics) < 1 {
		topics = append(topics, AllKey)
	}

	for _, topic := range topics {
		ShowDescription(topic)
	}
}

func ShowDescription(topic string) {
	if topic == AllKey {
		ShowAllDescriptions()
		return
	}

	help, ok := docMap[topic]
	if !ok {
		PrintfShadow("Unknown command '%s' type `drive help all` for entire usage documentation", topic)
		ShowAllDescriptions()
	} else {
		description, documentation := help[0], help[1:]
		PrintfShadow("Name\n\t%s - %s\n", topic, description)
		if len(documentation) >= 1 {
			PrintfShadow("Description\n")
			for _, line := range documentation {
				segments := formatText(line)
				for _, segment := range segments {
					PrintfShadow("\t%s", segment)
				}
			}
			PrintfShadow("\n* For usage flags: \033[32m`drive %s -h`\033[00m\n\n", topic)
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}

func formatText(text string) []string {
	splits := strings.Split(text, "\n")

	pr := prettywords.PrettyRubric{
		Limit: 80,
		Body:  splits,
	}

	return pr.Format()
}

func PrintfShadow(fmt_ string, args ...interface{}) {
	FprintfShadow(os.Stdout, fmt_, args...)
}

func StdoutPrintf(fmt_ string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, fmt_, args...)
}

func FprintfShadow(f io.Writer, fmt_ string, args ...interface{}) {
	sprinted := fmt.Sprintf(fmt_, args...)
	splits := formatText(sprinted)
	for _, split := range splits {
		fmt.Fprintf(f, "%s\n", split)
	}
}
