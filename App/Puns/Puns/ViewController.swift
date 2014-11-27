//
//  ViewController.swift
//  Puns
//
//  Created by Niraj  on 11/27/14.
//  Copyright (c) 2014 Niraj Pant. All rights reserved.
//

import UIKit

class ViewController: UIViewController {
    @IBOutlet weak var punLabel: UILabel!
    
    let punBook = PunBook()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        punLabel.text = punBook.randomPun()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    @IBAction func showPun() {        
        punLabel.text = punBook.randomPun()
    }

}

