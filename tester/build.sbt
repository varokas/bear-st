name := "tester"
organization := "io.bigbears"
version := "0.1.0-SNAPSHOT"

scalaVersion := "2.12.10"

libraryDependencies += "io.gatling.highcharts" % "gatling-charts-highcharts" % "3.3.1"
// libraryDependencies += "org.scalatest" %% "scalatest" % "3.0.3"
libraryDependencies += "io.gatling" % "gatling-test-framework" % "3.3.1"


enablePlugins(GatlingPlugin)
