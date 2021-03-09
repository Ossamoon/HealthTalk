//
//  SignInViewModel.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/07.
//

import Foundation
import Combine

class SignInViewModel : ObservableObject {
    @Published var inputName: String = ""
    @Published var inputPassword: String = ""
    
    private var url: URL! = URL(string: "http://18.236.209.128:8080/login")!
    private var httpMethod: String = "POST"
    private var cancellationToken: AnyCancellable?
    
    func setToken(token: Token) {
        UserDefaults.standard.set(token.token, forKey: "apiToken")
        let auth = Auth.shared
        auth.token = token.token
    }
    
    func run() {
        var request = URLRequest(url: self.url)
        request.httpMethod = self.httpMethod
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        let params: [String: String] = [
            "name": self.inputName,
            "password": self.inputPassword,
        ]
        request.httpBody = try! JSONSerialization.data(withJSONObject: params, options: [])
        print ("Request: \(request).")
        cancellationToken = URLSession.shared
            .dataTaskPublisher(for: request)
            .tryMap() { element -> Data in
                guard let httpResponse = element.response as? HTTPURLResponse, httpResponse.statusCode == 200 else {
                    throw URLError(.badServerResponse)
                }
                return element.data
            }
            .decode(type: Token.self, decoder: JSONDecoder())
            .sink(receiveCompletion: {
                print ("Received completion: \($0).")
            }, receiveValue: { token in
                print ("Received token: \(token).")
                self.setToken(token: token)
            })
    }
}
