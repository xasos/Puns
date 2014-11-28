//
//  ViewController.swift
//  Puns
//
//  Created by Niraj  on 11/27/14.
//  Copyright (c) 2014 Niraj Pant. All rights reserved.
//

import UIKit
import Alamofire

class ViewController: UIViewController {
    @IBOutlet weak var punLabel: UILabel!
    @IBOutlet weak var punButton: UIButton!
    
    let punBook = PunBook()
    let colorWheel = ColorWheel()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        punLabel.text = punBook.randomPun()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showPun() {
        var randomColor = colorWheel.randomColor()
        view.backgroundColor = randomColor
        punLabel.text = punBook.randomPun()
        punButton.tintColor = randomColor
        loadStuff()
    }
    
    func loadStuff() {
        Alamofire.request(.GET, "hhttp://httpbin.org/get", parameters: nil)
            .response { (request, response, data, error) in
                println(request)
                println(response)
                println(error)
        }
    }
}

