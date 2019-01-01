// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tPost struct {}
var Post tPost


func (_ tPost) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Index", args).URL
}

func (_ tPost) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Create", args).URL
}

func (_ tPost) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.Delete", args).URL
}

func (_ tPost) RedirectToPosts(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Post.RedirectToPosts", args).URL
}


type tSport struct {}
var Sport tSport


func (_ tSport) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sport.Index", args).URL
}

func (_ tSport) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sport.Create", args).URL
}

func (_ tSport) Delete(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sport.Delete", args).URL
}

func (_ tSport) RedirectToSports(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sport.RedirectToSports", args).URL
}


