//
//  SignUpView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/02/28.
//

import SwiftUI

struct SignUpView: View {
    
    @State var inputName: String = ""
    @State var inputPassword: String = ""
    @State var inputEmail: String = ""
    
    var body: some View {
        VStack {
            Text("新規ユーザー登録")
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
                TextField("メールアドレス", text: $inputEmail)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding()
            }.padding()
            
            Button(action: {
                print("SignUp処理")
                UserDefaults.standard.set("token", forKey: "apiToken")
                let auth = Auth.shared
                auth.token = "token"
            }) {
                Text("新規登録")
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

struct SignUpView_Previews: PreviewProvider {
    static var previews: some View {
        SignUpView()
    }
}
