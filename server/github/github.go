package github

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/ejholmes/buble"
	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/util"
)

// ResponseWriter wraps a buble.ResponseWriter.
type ResponseWriter interface {
	buble.ResponseWriter
}

// Response wraps a buble.Response.
type Response struct {
	buble.ResponseWriter
}

// Request wraps a buble.Request.
type Request struct {
	*buble.Request
}

// Authentic returns true if the calculated signature matches the
// signature from the request headers.
func (r *Request) Authentic(secret string) bool {
	sig, err := r.signature(secret)
	if err != nil {
		panic(err)
	}
	return r.Header.Get(SigHeader) == "sha1="+sig
}

// signature calculates the HMAC signature of the request body using the secret.
func (r *Request) signature(secret string) (string, error) {
	raw, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", nil
	}

	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(raw)
	return fmt.Sprintf("%x", mac.Sum(nil)), nil
}

// HandlerFunc defines the method signature for event handlers.
type HandlerFunc func(Shipr, ResponseWriter, *Request)

const (
	// EventHeader is the name of the header that determines what type of event this is.
	EventHeader = "X-GitHub-Event"

	// SigHeader is the name of the header that contains the sha1 of the payload.
	SigHeader = "X-Hub-Signature"
)

// Shipr is an interface that declares the methods we use from shipr.
type Shipr interface {
	Deploy(shipr.Description) error
	Notify(shipr.Notification) error
}

// GitHub demuxes incoming webhooks from GitHub and handles them.
type GitHub struct {
	shipr  Shipr
	router *mux.Router
	secret string
}

// New returns a new Handler.
func New(sh Shipr, secret string) *GitHub {
	h := &GitHub{
		shipr:  sh,
		router: mux.NewRouter(),
		secret: secret,
	}

	// Handlers
	h.Handle("deployment", Deployment)
	h.Handle("deployment_status", DeploymentStatus)

	return h
}

// Handle maps an event to a HandlerFunc.
func (g *GitHub) Handle(event string, fn HandlerFunc) {
	h := &buble.Handler{
		HandlerFunc: func(w buble.ResponseWriter, r *buble.Request) {
			resp := &Response{ResponseWriter: w}
			req := &Request{Request: r}

			// Ensure that the provided signature matches the calculated signature.
			if !req.Authentic(g.secret) {
				w.WriteHeader(403)
				return
			}

			fn(g.shipr, resp, req)
		},
	}
	g.router.Methods("POST").Headers(EventHeader, event).Handler(h)
}

func (g *GitHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.router.ServeHTTP(w, r)
}

// Deployment handles "deployment" events.
func Deployment(sh Shipr, w ResponseWriter, r *Request) {
	var d github.Deployment
	r.Decode(&d)

	err := sh.Deploy(newDeployment(&d))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

// DeploymentStatus handles "deployment_status" events.
func DeploymentStatus(sh Shipr, w ResponseWriter, r *Request) {
	var d github.DeploymentStatus
	r.Decode(&d)

	err := sh.Notify(newDeploymentStatus(&d))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

// deployment wraps a github.Deployment to implement the shipr.Description interface.
type deployment struct {
	GitHubDeployment *github.Deployment
}

// NewDeployment returns a new Deployment.
func newDeployment(d *github.Deployment) *deployment {
	return &deployment{GitHubDeployment: d}
}

func (d *deployment) Guid() int {
	return *d.GitHubDeployment.ID
}

func (d *deployment) RepoName() shipr.RepoName {
	return shipr.RepoName(util.SafeString(d.GitHubDeployment.Repository.FullName))
}

func (d *deployment) Sha() string {
	return util.SafeString(d.GitHubDeployment.Sha)
}

func (d *deployment) Ref() string {
	return util.SafeString(d.GitHubDeployment.Ref)
}

func (d *deployment) Environment() string {
	return util.SafeString(d.GitHubDeployment.Environment)
}

func (d *deployment) Description() string {
	return util.SafeString(d.GitHubDeployment.Description)
}

// deploymentStatus wraps a github.DeploymentStatus
type deploymentStatus struct {
	*deployment
	GitHubDeploymentStatus *github.DeploymentStatus
}

// newDeploymentStatus returns a new DeploymentStatus.
func newDeploymentStatus(d *github.DeploymentStatus) *deploymentStatus {
	return &deploymentStatus{
		deployment:             newDeployment(d.Deployment),
		GitHubDeploymentStatus: d,
	}
}

func (n *deploymentStatus) RepoName() shipr.RepoName {
	return shipr.RepoName("TODO")
}

func (n *deploymentStatus) URL() *url.URL {
	u, err := url.Parse("http://www.google.com")
	if err != nil {
		panic(err)
	}
	return u
}

func (n *deploymentStatus) User() string {
	return "TODO"
}

func (n *deploymentStatus) State() string {
	return util.SafeString(n.GitHubDeploymentStatus.State)
}
