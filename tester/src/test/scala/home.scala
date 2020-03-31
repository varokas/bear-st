package tester

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class HomePageSimulation extends Simulation {

  val httpProtocol = http.baseUrl("https://localhost:3000")

  val numberOfUsers = 10
  val duration = 10

  val scn = scenario("Home")
    .exec(http("open_page")
      .get("/")
      .check(status.is(200)))
    .inject(constantUsersPerSec(numberOfUsers) during (duration seconds))
    .throttle(reachRps(numberOfUsers.toInt) in (5 seconds), holdFor(duration seconds))

  setUp(scn).protocols(httpProtocol)
}