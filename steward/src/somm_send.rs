use abscissa_core::tracing::info;
use deep_space::coin::Coin;
use deep_space::error::CosmosGrpcError;
use deep_space::Contact;
use deep_space::Fee;
use deep_space::Msg;
use deep_space::PrivateKey as CosmosPrivateKey;
use gravity_bridge::gravity_proto::cosmos_sdk_proto::cosmos::{
    base::abci::v1beta1::TxResponse, tx::v1beta1::BroadcastMode,
};
use somm_proto::cork::Cork;
use somm_proto::cork::{MsgScheduleCorkRequest, MsgSubmitCorkRequest};
use std::{result::Result, time::Duration};

pub const TIMEOUT: Duration = Duration::from_secs(60);
pub const MEMO: &str = "Sent using Somm Orchestrator";

pub async fn send_cork(
    contact: &Contact,
    cork: Cork,
    delegate_address: String,
    delegate_key: &CosmosPrivateKey,
    fee: Coin,
) -> Result<TxResponse, CosmosGrpcError> {
    let msg = MsgSubmitCorkRequest {
        cork: Some(cork),
        signer: delegate_address,
    };
    let msg = Msg::new("/cork.v1.MsgSubmitCorkRequest", msg);
    __send_messages(contact, delegate_key, fee, vec![msg]).await
}

pub async fn schedule_cork(
    contact: &Contact,
    cork: Cork,
    delegate_address: String,
    delegate_key: &CosmosPrivateKey,
    fee: Coin,
    block_height: u64,
) -> Result<TxResponse, CosmosGrpcError> {
    let msg = MsgScheduleCorkRequest {
        cork: Some(cork),
        signer: delegate_address,
        block_height,
    };
    let msg = Msg::new("/cork.v1.MsgScheduleCorkRequest", msg);
    __send_messages(contact, delegate_key, fee, vec![msg]).await
}

async fn __send_messages(
    contact: &Contact,
    cosmos_key: &CosmosPrivateKey,
    fee: Coin,
    messages: Vec<Msg>,
) -> Result<TxResponse, CosmosGrpcError> {
    let cosmos_address = cosmos_key.to_address(&contact.get_prefix())?;

    let fee = Fee {
        amount: vec![fee],
        gas_limit: 500_000u64 * (messages.len() as u64),
        granter: None,
        payer: None,
    };

    let args = contact.get_message_args(cosmos_address, fee).await?;

    let msg_bytes = cosmos_key.sign_std_msg(&messages, args, MEMO)?;

    let response = contact
        .send_transaction(msg_bytes, BroadcastMode::Sync)
        .await?;

    info!("SubmitCork response: {:?}", response);

    contact.wait_for_tx(response, TIMEOUT).await
}
