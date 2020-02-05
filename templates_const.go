package main

// template assets

const FileSystemAssetDir = "assets/"

const (
	// TODO 1092387750 write a generate to set this

	TemplatesHTMLPrefix AssetPath = "/templates/html/"

	BaseHtml       AssetPath = TemplatesHTMLPrefix + "base.html"
	BaseHeaderHtml AssetPath = TemplatesHTMLPrefix + "base_header.html"
	BaseFooterHtml AssetPath = TemplatesHTMLPrefix + "base_footer.html"

	IndexHtml    AssetPath = TemplatesHTMLPrefix + "index.html"
	Error404Html AssetPath = TemplatesHTMLPrefix + "404.html"
	Error5XXHtml AssetPath = TemplatesHTMLPrefix + "5xx.html"

	LoginHtml  AssetPath = TemplatesHTMLPrefix + "login.html"
	SignUpHtml AssetPath = TemplatesHTMLPrefix + "signup.html"

	// logged in users only
	DashboardHtml AssetPath = TemplatesHTMLPrefix + "dashboard.html"

	// admin users only
	AdminDashboardHtml AssetPath = TemplatesHTMLPrefix + "admin_dashboard.html"

// super admin users only

)

// Content Keys
const (
	UsersContentKey   ContentKey = "Users"   // slice of *AuthLayerUsers
	BooksContentKey   ContentKey = "Books"   // a slice of books
	AuthorsContentKey ContentKey = "Authors" // a slice of authors
)

// Content Error Keys
const (
	InvalidPasswordContentErrorKey ContentErrorKey = "InvalidPassword"
	UnknownErrorContentErrorKey    ContentErrorKey = "UnknownError"
)
