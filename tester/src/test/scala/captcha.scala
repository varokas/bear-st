package tester

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import scala.concurrent.duration._

class CaptchaPageSimulation extends Simulation {

  val httpProtocol = http.baseUrl("http://localhost:3000")

  val numberOfUsers = 100
  val duration = 2

  val scn = scenario("Captcha")
    .exec(http("open_page")
      .get("/")
      .check(status.is(200)))
    .inject(constantUsersPerSec(numberOfUsers) during (duration seconds))
    .throttle(reachRps(numberOfUsers.toInt) in (5 seconds), holdFor(duration seconds))

  setUp(scn).protocols(httpProtocol)
}