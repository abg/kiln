package main

import (
	"log"
	"os"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/kiln/builder"
	"github.com/pivotal-cf/kiln/commands"
	"github.com/pivotal-cf/kiln/helper"
	"github.com/pivotal-cf/kiln/internal/baking"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	var global struct {
		Help bool `short:"h" long:"help" description:"prints this usage information" default:"false"`
	}

	args, err := jhanda.Parse(&global, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	globalFlagsUsage, err := jhanda.PrintUsage(global)
	if err != nil {
		log.Fatal(err)
	}

	filesystem := helper.NewFilesystem()
	zipper := builder.NewZipper()
	handcraftReader := builder.NewMetadataReader(filesystem, logger)
	iconEncoder := builder.NewIconEncoder(filesystem)
	variablesDirectoryReader := builder.NewMetadataPartsDirectoryReaderWithTopLevelKey("variables")
	metadataBuilder := builder.NewMetadataBuilder(
		variablesDirectoryReader,
		handcraftReader,
		logger,
		iconEncoder,
	)
	md5SumCalculator := helper.NewFileMD5SumCalculator()
	interpolator := builder.NewInterpolator()
	tileWriter := builder.NewTileWriter(filesystem, &zipper, logger, md5SumCalculator)

	templateVariablesParser := baking.NewTemplateVariablesService()

	releaseManifestReader := builder.NewReleaseManifestReader()
	releasesService := baking.NewReleasesService(releaseManifestReader)

	stemcellManifestReader := builder.NewStemcellManifestReader(filesystem)
	stemcellService := baking.NewStemcellService(logger, stemcellManifestReader)

	formDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	formsService := baking.NewFormsService(logger, formDirectoryReader)

	instanceGroupDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	instanceGroupsService := baking.NewInstanceGroupsService(logger, instanceGroupDirectoryReader)

	jobsDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	jobsService := baking.NewJobsService(logger, jobsDirectoryReader)

	propertiesDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	propertiesService := baking.NewPropertiesService(logger, propertiesDirectoryReader)

	runtimeConfigsDirectoryReader := builder.NewMetadataPartsDirectoryReader()
	runtimeConfigsService := baking.NewRuntimeConfigsService(logger, runtimeConfigsDirectoryReader)

	commandSet := jhanda.CommandSet{}
	commandSet["help"] = commands.NewHelp(os.Stdout, globalFlagsUsage, commandSet)
	commandSet["bake"] = commands.NewBake(
		metadataBuilder,
		interpolator,
		tileWriter,
		logger,
		yaml.Marshal,
		templateVariablesParser,
		releasesService,
		stemcellService,
		formsService,
		instanceGroupsService,
		jobsService,
		propertiesService,
		runtimeConfigsService,
	)

	var command string
	if len(args) > 0 {
		command, args = args[0], args[1:]
	}

	if global.Help {
		command = "help"
	}

	if command == "" {
		command = "help"
	}

	err = commandSet.Execute(command, args)
	if err != nil {
		log.Fatal(err)
	}
}
