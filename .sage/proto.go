package main

import (
	"context"
	"os"

	"go.einride.tech/sage/sg"
	"go.einride.tech/sage/sgtool"
	"go.einride.tech/sage/tools/sgbuf"
)

type Proto sg.Namespace

func (Proto) All(ctx context.Context) error {
	sg.Deps(ctx, Proto.BufFormat, Proto.BufLint, Proto.BufGenerate)
	sg.Deps(ctx, Proto.BufGenerate)
	return nil
}

func (Proto) BufLint(ctx context.Context) error {
	sg.Logger(ctx).Println("linting proto files...")
	cmd := sgbuf.Command(ctx, "lint")
	cmd.Dir = sg.FromGitRoot("internal", "examples", "proto")
	return cmd.Run()
}

func (Proto) BufFormat(ctx context.Context) error {
	sg.Logger(ctx).Println("formatting proto files...")
	cmd := sgbuf.Command(ctx, "format", "--write")
	cmd.Dir = sg.FromGitRoot("internal", "examples", "proto")
	return cmd.Run()
}

func (Proto) ProtocGenGo(ctx context.Context) error {
	sg.Logger(ctx).Println("installing...")
	_, err := sgtool.GoInstallWithModfile(ctx, "google.golang.org/protobuf/cmd/protoc-gen-go", sg.FromGitRoot("go.mod"))
	return err
}

func (Proto) ProtocGenBigqueryJSONSchema(ctx context.Context) error {
	sg.Logger(ctx).Println("building...")
	return sg.Command(
		ctx,
		"go",
		"build",
		"-o",
		sg.FromBinDir("protoc-gen-bq-json-schema"),
		"./protoc-gen-bq-json-schema",
	).Run()
}

func (Proto) BufGenerate(ctx context.Context) error {
	sg.Deps(ctx, Proto.ProtocGenGo, Proto.ProtocGenBigqueryJSONSchema)
	sg.Logger(ctx).Println("generating example proto stubs...")
	if err := os.RemoveAll(sg.FromGitRoot("internal", "examples", "proto", "gen")); err != nil {
		return err
	}
	cmd := sgbuf.Command(ctx, "generate", "--template", "buf.gen.yaml", "--path", "einride")
	cmd.Dir = sg.FromGitRoot("internal", "examples", "proto")
	return cmd.Run()
}
