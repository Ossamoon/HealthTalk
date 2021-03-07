//
//  Common.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import Foundation

class Auth : ObservableObject {
    @Published var token: String? = nil
    static let shared = Auth()
    private init() {
    }
}
