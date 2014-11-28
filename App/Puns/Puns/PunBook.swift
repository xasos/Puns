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
    func getRandomPun() -> String {
        Alamofire.request(.GET, "http://getpuns.herokuapp.com/api/random", parameters: nil)
            .responseJSON { (request, response, json, error) in
                println(response)
                println(json)
                if (json != nil) {
                    var jsonObj = JSON(json!)
                    let data = jsonObj["pun"].stringValue
                }
                else {
                    println("error")
                }
        }
        return "asdasd"
    }
}

