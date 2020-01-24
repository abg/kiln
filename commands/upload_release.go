package commands

import (
	"fmt"
	"github.com/pivotal-cf/kiln/fetcher"
	"github.com/pivotal-cf/kiln/internal/cargo"
	"github.com/pivotal-cf/kiln/release"
	"log"

	"github.com/pivotal-cf/kiln/builder"

	"github.com/pivotal-cf/jhanda"
	"gopkg.in/src-d/go-billy.v4"
)

type UploadRelease struct {
	FS                     billy.Filesystem
	KilnfileLoader         KilnfileLoader
	ReleaseUploaderFactory ReleaseUploaderFactory
	Logger                 *log.Logger

	Options struct {
		Kilnfile       string   `short:"kf" long:"kilnfile" default:"Kilnfile" description:"path to Kilnfile"`
		Variables      []string `short:"vr" long:"variable" description:"variable in key=value format"`
		VariablesFiles []string `short:"vf" long:"variables-file" description:"path to variables file"`

		ReleaseSource string `short:"rs" long:"release-source" required:"true" description:"name of the release source specified in the Kilnfile"`
		LocalPath     string `short:"lp" long:"local-path" required:"true" description:"path to BOSH release tarball"`
	}
}

//go:generate counterfeiter -o ./fakes/release_uploader_factory.go --fake-name ReleaseUploaderFactory . ReleaseUploaderFactory
type ReleaseUploaderFactory interface {
	ReleaseUploader(sourceID string, kilnfile cargo.Kilnfile) (fetcher.ReleaseUploader, error)
}

func (command UploadRelease) Execute(args []string) error {
	_, err := jhanda.Parse(&command.Options, args)
	if err != nil {
		return err
	}

	kilnfile, _, err := command.KilnfileLoader.LoadKilnfiles(
		command.FS,
		command.Options.Kilnfile,
		command.Options.VariablesFiles,
		command.Options.Variables,
	)
	if err != nil {
		return fmt.Errorf("error loading Kilnfiles: %w", err)
	}

	releaseSource, err := command.ReleaseUploaderFactory.ReleaseUploader(command.Options.ReleaseSource, kilnfile)
	if err != nil {
		return fmt.Errorf("error finding release source: %w", err)
	}

	file, err := command.FS.Open(command.Options.LocalPath)
	if err != nil {
		return fmt.Errorf("could not open release: %w", err)
	}

	manifestReader := builder.NewReleaseManifestReader(command.FS)
	part, err := manifestReader.Read(command.Options.LocalPath)
	if err != nil {
		return fmt.Errorf("error reading the release manifest: %w", err)
	}

	manifest := part.Metadata.(builder.ReleaseManifest)

	requirement := release.Requirement{Name: manifest.Name, Version: manifest.Version}
	_, found, err := releaseSource.GetMatchedRelease(requirement)
	if err != nil {
		return fmt.Errorf("couldn't query release source: %w", err)
	}

	if found {
		return fmt.Errorf("a release with name %q and version %q already exists on %s",
			manifest.Name, manifest.Version, command.Options.ReleaseSource)
	}

	err = releaseSource.UploadRelease(manifest.Name, manifest.Version, file)
	if err != nil {
		return fmt.Errorf("error uploading the release: %w", err)
	}

	command.Logger.Println("Upload succeeded")

	return nil
}

func (command UploadRelease) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "Uploads a BOSH Release to an S3 release source for use in kiln fetch",
		ShortDescription: "uploads a BOSH release to an s3 release_source",
		Flags:            command.Options,
	}
}