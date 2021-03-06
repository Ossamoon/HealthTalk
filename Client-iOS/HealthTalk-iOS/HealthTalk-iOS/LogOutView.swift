//
//  LogOutView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct LogOutView: View {
    var body: some View {
        VStack(spacing: 80) {
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

struct LogOutView_Previews: PreviewProvider {
    static var previews: some View {
        LogOutView()
    }
}
