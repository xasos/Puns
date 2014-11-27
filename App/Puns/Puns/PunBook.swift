//
//  PunBook.swift
//  Puns
//
//  Created by Niraj  on 11/27/14.
//  Copyright (c) 2014 Niraj Pant. All rights reserved.
//

import Foundation

struct PunBook {
    let punsArray = [
        "asdasd",
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
