package controllers

import javax.inject._
import play.api._
import play.api.libs.json._
import play.api.mvc._


class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  private val products = List("Product1", "Product2", "Product3", "Product3")

  def getAll(): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(products))
  }
}
