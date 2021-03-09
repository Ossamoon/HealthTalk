//
//  SignInView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct SignInView: View {
    @ObservedObject var viewModel = SignInViewModel()
    
    var body: some View {
        VStack {
            Text("ログイン画面")
                .font(.title)
                .fontWeight(.bold)
                .padding()
            
            VStack {
                TextField("ユーザー名", text: $viewModel.inputName)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding()
                SecureField("パスワード", text: $viewModel.inputPassword)
                    .textFieldStyle(RoundedBorderTextFieldStyle())
                    .padding()
            }.padding()
            
            Button(action: {
                viewModel.run()
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
