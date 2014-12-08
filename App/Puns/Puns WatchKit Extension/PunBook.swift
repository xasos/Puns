//
//  PunBook.swift
//  Puns
//
//  Created by Niraj  on 11/27/14.
//  Copyright (c) 2014 Niraj Pant. All rights reserved.
//

import Foundation
import Alamofire

class PunBook {
    var punString: String = ""
    func getRandomPun() -> String {
        Alamofire.request(.GET, "http://getpuns.herokuapp.com/api/random", parameters: nil)
            .responseJSON { (request, response, json, error) in
                if (json != nil) {
                    var jsonObj = JSON(json!)
                    var data = jsonObj["Pun"].stringValue
                    self.punString = data
                }
                else {
                    println("error")
                }
        }
        return punString
    }
}