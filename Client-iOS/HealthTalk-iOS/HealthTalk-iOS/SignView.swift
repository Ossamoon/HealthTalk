//
//  SignView.swift
//  HealthTalk-iOS
//
//  Created by 齋藤修 on 2021/03/06.
//

import SwiftUI

struct SignView: View {
    var body: some View {
           NavigationView {
               VStack(spacing: 60) {
                   NavigationLink(destination: SignUpView()) {
                       Text("新規登録")
                   }
                   NavigationLink(destination: SignInView()) {
                       Text("ログイン")
                   }
               }
           }
       }
}

struct SignView_Previews: PreviewProvider {
    static var previews: some View {
        SignView()
    }
}
