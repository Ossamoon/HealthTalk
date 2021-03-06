//
//  SignInView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct SignInView: View {
    @State var inputName: String = ""
    @State var inputPassword: String = ""
    
    var body: some View {
        VStack {
            Text("ログイン画面")
                .font(.title)
                .fontWeight(.bold)
                .padding()
            
            VStack {
                TextField("ユーザー名", text: $inputName)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding()
                SecureField("パスワード", text: $inputPassword)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding()
            }.padding()
            
            Button(action: {
                print("SignIn処理")
                UserDefaults.standard.set("token", forKey: "apiToken")
                let auth = Auth.shared
                auth.token = "token"
            }) {
                Text("ログイン")
                    .font(.title2)
                    .fontWeight(.bold)
                    .foregroundColor(.white)
                    .padding()
                    .background(Color.accentColor)
                    .cornerRadius(8.0)
            }.padding()
        }
    }
}

struct SignInView_Previews: PreviewProvider {
    static var previews: some View {
        SignInView()
    }
}
