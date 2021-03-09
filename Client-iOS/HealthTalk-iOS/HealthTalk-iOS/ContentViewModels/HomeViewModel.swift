//
//  HomeViewModel.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/08.
//

import Foundation
import Combine

class HomeViewModel : ObservableObject {
    private var url: URL! = URL(string: "http://18.236.209.128:8080/api/user")!
    private var httpMethod: String = "GET"
    private var cancellationToken: AnyCancellable?
    
    func run() {
        var request = URLRequest(url: self.url)
        request.httpMethod = self.httpMethod
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        request.addValue("Bearer " + Auth.shared.token!, forHTTPHeaderField: "Authorization")
        print ("Request: \(request).")
        cancellationToken = URLSession.shared
            .dataTaskPublisher(for: request)
            .tryMap() { element -> Data in
                guard let httpResponse = element.response as? HTTPURLResponse, httpResponse.statusCode == 200 else {
                    throw URLError(.badServerResponse)
                }
                return element.data
            }
            .decode(type: User.self, decoder: JSONDecoder())
            .sink(receiveCompletion: {
                print ("Received completion: \($0).")
            }, receiveValue: { token in
                print ("Received token: \(token).")
            })
    }
}
