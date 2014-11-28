//
//  PunBook.swift
//  Puns
//
//  Created by Niraj  on 11/27/14.
//  Copyright (c) 2014 Niraj Pant. All rights reserved.
//

import Foundation
import Alamofire

struct PunBook {
    let punsArray = [
        "When a clock is hungry it goes back four seconds.",
        "asdasdasda",
        "adasdasdasdasd"
    ]
    
    func randomPun() -> String {
        var unsignedArrayCount = UInt32(punsArray.count)
        var unsignedRandomNumber = arc4random_uniform(unsignedArrayCount)
        var randomNumber = Int(unsignedRandomNumber)
        
        return punsArray[randomNumber]
    }
}

func getRandomPun() {
    Alamofire.request(.GET, "http://getpuns.herokuapp.com/api/random", parameters: nil)
        .response { (request, response, data, error) in
            println(request)
            println(response)
            println(error)
    }
}
