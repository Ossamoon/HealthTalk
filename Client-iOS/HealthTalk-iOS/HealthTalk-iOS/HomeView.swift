//
//  HomeView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct HomeView: View {
    var body: some View {
        VStack {
            Text("HomeView!")
            Button(action: {
                print("LogOut処理")
                UserDefaults.standard.set(nil, forKey: "apiToken")
                let auth = Auth.shared
                auth.token = nil
            }) {
                Text("ログアウトする")
            }.padding()
        }
    }
}

struct HomeView_Previews: PreviewProvider {
    static var previews: some View {
        HomeView()
    }
}
