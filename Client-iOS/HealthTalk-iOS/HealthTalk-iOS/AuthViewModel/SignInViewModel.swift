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
    
    private var url: URL! = URL(string: "https://34.220.52.80/endpoint")!
    private var httpMethod: String = "POST"
    private var subscriptions = Set<AnyCancellable>()
    private var fetchedToken: Data?
    private var response: URLResponse?
    
    func run() {
        var request = URLRequest(url: self.url)
        request.httpMethod = self.httpMethod
        URLSession.shared
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
                    UserDefaults.standard.set("token", forKey: "apiToken")
                    let auth = Auth.shared
                    auth.token = "token"
                })
            .store(in: &subscriptions)
    }
}
